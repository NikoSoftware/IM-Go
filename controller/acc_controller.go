package controller

import (
	"IM-Go/common"
	"IM-Go/wsocket"
	"fmt"
	"time"
)

func LoginController(client *wsocket.Client, seq int, message interface{}) (code uint32, msg string, data interface{}) {

	code = common.OK
	currentTime := time.Now().UnixMilli()

	fmt.Println(currentTime)

	return
}
