package models

import (
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
	processingTime, _ := strconv.ParseFloat(ddbItem["processing_time"].String(), 64)
	chewiness, _ := strconv.Atoi(ddbItem["chewiness"].String())
	return Item{
		Id:             ddbItem["id"].String(),
		Timestamp:      ddbItem["timestamp"].String(),
		Chewiness:      chewiness,
		ProcessingTime: processingTime,
		Data:           ddbItem["data"].String(),
	}
}
