package worker

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/bvtujo/copilot-overview/frontend"
	"math"
	"math/rand"
	"os"
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
			QueueUrl:            aws.String(os.Getenv("COPILOT")),
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

func (l *Listener) processMessage(message *sqs.Message) {
	var msg frontend.Message
	if err := json.Unmarshal([]byte(aws.StringValue(message.Body)), &msg); err != nil {
		fmt.Println(fmt.Errorf("process message with id %s: %w", aws.StringValue(message.MessageId), err))
		return
	}
	// Chew on message.
	sleepTime := rand.NormFloat64()*math.Sqrt(float64(msg.Chewiness)) + float64(msg.Chewiness)
	time.Sleep(time.Duration(float64(time.Second) * sleepTime))

	_, err := l.ddb.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": &dynamodb.AttributeValue{
				S: aws.String(msg.Id),
			},
			"timestamp": &dynamodb.AttributeValue{
				S: aws.String(fmt.Sprintf("%v", time.Now().Unix())),
			},
			"processing_time": &dynamodb.AttributeValue{
				N: aws.String(fmt.Sprintf("%v", sleepTime)),
			},
			"data": &dynamodb.AttributeValue{
				S: aws.String(msg.Data),
			},
		},
		TableName: aws.String(l.table),
	})
	if err != nil {
		fmt.Println(fmt.Errorf("put ddb item with id %s and data %q: %w", msg.Id, msg.Data, err))
		return
	}
}
