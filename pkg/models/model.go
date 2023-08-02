package models

type Message struct {
	Id        string `json:"id"`
	Chewiness int    `json:"chewiness"`
	Data      string `json:"data"`
}
