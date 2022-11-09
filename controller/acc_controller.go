package controller

import (
	"IM-Go/common"
	"IM-Go/db"
	"IM-Go/model"
	"IM-Go/wsocket"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func LoginController(client *wsocket.Client, seq int, message []byte) (code uint32, msg string, data interface{}) {

	code = common.OK
	currentTime := time.Now().UnixMilli()

	fmt.Println(currentTime)

	loginReq := &model.User{}

	err := json.Unmarshal(message, loginReq)
	if err != nil {
		log.Printf("登陆数据解析异常:%v", err)
		code = common.ParameterIllegal
		return
	}

	user := &model.User{}

	//	db.Db.Debug().Model(&model.Course{}).Preload("Teacher").Find(&course)
	//db.Db.Debug().Where().Model(&model.User{}).Preload("User").First(user)
	db.Db.Debug().First(user, loginReq.Id)

	//未查到此数据
	if user.Id == 0 {
		code = common.UnauthorizedUserId
		return

	}

	if user.Password != loginReq.Password {
		code = common.PasswordErr
		return
	}

	//加入登陆用户
	wsocket.DisposeChanService.Login <- &wsocket.Login{
		UserId: user.Id,
		Client: client,
	}

	return
}
