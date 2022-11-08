package wsocket

import (
	"IM-Go/common"
	"IM-Go/model"
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

type DisposeFunc func(client *Client, seq int, message interface{}) (code uint32, msg string, data interface{})

var (
	handlers        = make(map[string]DisposeFunc)
	handlersRWMutex sync.RWMutex
)

func Register(key string, method DisposeFunc) {
	defer handlersRWMutex.Unlock()

	handlersRWMutex.Lock()
	handlers[key] = method

}

func GetHandle(key string) (value DisposeFunc, ok bool) {

	defer handlersRWMutex.Unlock()
	handlersRWMutex.Lock()

	value, ok = handlers[key]

	return

}

func MsgProcess(client *Client, message []byte) {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("MsgProcess处理异常 addr: %v", client.Addr)
		}
	}()

	request := &model.Message{}
	err := json.Unmarshal(message, request)
	if err != nil {
		fmt.Println(err)
		client.SendMsg([]byte("消息不合法"))
	}

	var (
		code uint32
		msg  string
		data interface{}
	)

	if value, ok := GetHandle(request.Cmd); ok {
		code, msg, data = value(client, request.Id, request.Data)
	} else {
		code = common.RoutingNotExist
		msg = "参数错误，无此操作"
	}

	msg = common.GetErrorMessage(code, msg)

	responseHead := model.NewResponseHead(request.Id, request.Cmd, code, msg, data)

	log.Printf("response消息：%v", responseHead)

	headByte, err := json.Marshal(responseHead)

	client.SendMsg(headByte)

}
