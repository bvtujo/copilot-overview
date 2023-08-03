package models

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"strconv"
)

type Message struct {
	Id        string `json:"id,omitempty"`
	Chewiness int    `json:"chewiness,omitempty"`
	Data      string `json:"data,omitempty"`
}

type Item struct {
	Id             string  `json:"id"`
	Timestamp      string  `json:"timestamp"`
	Chewiness      int     `json:"chewiness"`
	ProcessingTime float64 `json:"processing_time"`
	Data           string  `json:"data"`
}

func NewItemFromDDB(ddbItem map[string]*dynamodb.AttributeValue) Item {
	processingTime, _ := strconv.ParseFloat(aws.StringValue(ddbItem["processing_time"].N), 64)
	chewiness, _ := strconv.Atoi(aws.StringValue(ddbItem["chewiness"].N))
	return Item{
		Id:             aws.StringValue(ddbItem["id"].S),
		Timestamp:      aws.StringValue(ddbItem["timestamp"].S),
		Chewiness:      chewiness,
		ProcessingTime: processingTime,
		Data:           aws.StringValue(ddbItem["data"].S),
	}
}
