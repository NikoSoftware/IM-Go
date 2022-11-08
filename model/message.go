package model

type Message struct {
	Id   int         `json:"id""`
	Cmd  string      `json:"cmd""`
	Data interface{} `json:"data"`
}
