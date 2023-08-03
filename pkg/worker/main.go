package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/bvtujo/copilot-overview/pkg/models"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Listener struct {
	client sqs.SQS
	ddb    dynamodb.DynamoDB

	queue string
	table string
}

func main() {
	sess := session.Must(session.NewSession())
	client := sqs.New(sess)
	ddb := dynamodb.New(sess)
	queue := os.Getenv("COPILOT_QUEUE_URI")

	table := os.Getenv("DB_NAME")
	mp := &Listener{*client, *ddb, queue, table}
	mp.listen()
}

func (l *Listener) listen() {
	for {
		fmt.Println("polling for a message; wait 10 seconds")
		resp, err := l.client.ReceiveMessage(&sqs.ReceiveMessageInput{
			MaxNumberOfMessages: aws.Int64(1),
			QueueUrl:            aws.String(l.queue),
			VisibilityTimeout:   aws.Int64(15),
			WaitTimeSeconds:     aws.Int64(10),
		})
		if len(resp.Messages) < 1 {
			fmt.Println("no messages. sleeping 1s")
			time.Sleep(time.Second)
		}
		if err != nil {
			fmt.Println(fmt.Errorf("receive message: %w", err))
			continue
			time.Sleep(time.Second)
		}
		for _, message := range resp.Messages {
			go l.processMessage(message)
		}
	}
}

type SQSMessage struct {
	Body SNSMessage `json:"Body"`
}

type SNSMessage struct {
	Type      string `json:"Type,omitempty"`
	MessageId string `json:"MessageId,omitempty"`
	Message   string `json:"Message,omitempty"`
}

func (l *Listener) processMessage(message *sqs.Message) {
	var temp SNSMessage
	if err := json.Unmarshal([]byte(aws.StringValue(message.Body)), &temp); err != nil {
		fmt.Println("message", aws.StringValue(message.Body))
		fmt.Println(fmt.Errorf("process message with id %s: %w", aws.StringValue(message.MessageId), err))
		return
	}
	data := temp.Message
	data = strings.ReplaceAll(data, `\"`, `"`)
	var msg models.Message
	err := json.Unmarshal([]byte(data), &msg)
	if err != nil {
		fmt.Println(fmt.Errorf("unmarshal message %v: %w", data, err))
	}
	// Chew on message.
	sleepTime := rand.NormFloat64()*math.Sqrt(float64(msg.Chewiness)) + float64(msg.Chewiness)
	time.Sleep(time.Duration(float64(time.Second) * sleepTime))
	fmt.Println(fmt.Sprintf("inserting item after %d seconds: %"))

	ddbItem := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(msg.Id),
			},
			"timestamp": {
				S: aws.String(fmt.Sprintf("%v", time.Now().Unix())),
			},
			"processing_time": {
				N: aws.String(fmt.Sprintf("%v", sleepTime)),
			},
			"chewiness": {
				N: aws.String(fmt.Sprintf("%v", msg.Chewiness)),
			},
			"data": {
				S: aws.String(msg.Data),
			},
		},
		TableName: aws.String(l.table),
	}
	_, err = l.ddb.PutItem(ddbItem)
	if err != nil {
		fmt.Println(fmt.Errorf("put ddb item with id %s and data %q: %w", msg.Id, msg.Data, err))
		return
	}
	l.client.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(l.queue),
		ReceiptHandle: message.ReceiptHandle,
	})
	fmt.Printf("processed message in %v seconds\n", sleepTime)
}
