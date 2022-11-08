package wsocket

import (
	"IM-Go/model"
	"encoding/json"
	"fmt"
)

type DisposeFunc func(client *Client, seq string, message []byte) (id uint32, cmd string, data interface{})

func MsgProcess(client *Client, message []byte) {

	msg := &model.Message{}
	err := json.Unmarshal(message, msg)
	if err != nil {
		fmt.Println(err)
		client.SendMsg("err", []byte("消息不合法"))
	}

}
