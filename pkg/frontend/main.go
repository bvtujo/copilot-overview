package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/bvtujo/copilot-overview/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

func setUpSNS(sess *session.Session) (*sns.SNS, string, error) {
	snsClient := sns.New(sess)

	topicARNs := os.Getenv("COPILOT_SNS_TOPIC_ARNS")
	if topicARNs == "" {
		return nil, "", errors.New("can't find env var COPILOT_SNS_TOPIC_ARNS")
	}

	topicMap := make(map[string]string)
	if err := json.Unmarshal([]byte(topicARNs), &topicMap); err != nil {
		return nil, "", err
	}
	topicARN := topicMap["requests"]
	return snsClient, topicARN, nil
}

type DDB struct {
	client dynamodb.DynamoDB

	table string
}

func setUpDB(sess *session.Session) *DDB {
	ddbClient := dynamodb.New(sess)

	table := os.Getenv("DB_NAME")

	return &DDB{*ddbClient, table}
}

func (d *DDB) list() {
	d.client.Scan(&dynamodb.ScanInput{
		ExpressionAttributeNames:  nil,
		ExpressionAttributeValues: nil,
		FilterExpression:          nil,
		IndexName:                 nil,
		Limit:                     nil,
		ProjectionExpression:      nil,
		ReturnConsumedCapacity:    nil,
		ScanFilter:                nil,
		Segment:                   nil,
		Select:                    nil,
		TableName:                 aws.String(d.table),
		TotalSegments:             nil,
	})
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	mySession := session.Must(session.NewSession())

	snsClient, topicARN, err := setUpSNS(mySession)
	if err != nil {
		panic(err.Error())
	}
	r.GET("/healthcheck", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	//ddbClient, err := setUpDB()
	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		chewinessRaw := c.Query("chewiness")
		chewiness, err := strconv.Atoi(chewinessRaw)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		data := c.DefaultQuery("data", "")
		dataBytes, err := json.Marshal(&models.Message{id, chewiness, data})
		if err != nil {
			fmt.Println(fmt.Errorf("unmarshal data from json: %w", err))
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		_, err = snsClient.Publish(&sns.PublishInput{
			Message:  aws.String(string(dataBytes)),
			TopicArn: aws.String(topicARN),
		})
		if err != nil {
			fmt.Println(fmt.Errorf("error sending message: %w", err))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.Status(http.StatusOK)
	})

	r.GET("/list", func(c *gin.Context) {
		client := setUpDB()
	})

	r.Run(":8080")
}
