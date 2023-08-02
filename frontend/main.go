package frontend

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type Message struct {
	Id        string `json:"id"`
	Chewiness int    `json:"chewiness"`
	Data      string `json:"data"`
}

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

//func setUpDB(sess *session.Session) (*dynamodb.DynamoDB, string, error) {
//	ddbClient := dynamodb.New(sess)
//
//	return ddbClient,
//}

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
		dataBytes, err := json.Marshal(&Message{id, chewiness, data})
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

	//r.GET("/list", func(c *gin.Context) {
	//	client := setUpDB()
	//})

	r.Run(":8080")
}
