package routers

import (
	"IM-Go/controller"
	"IM-Go/wsocket"
)

func WSocketHandlerRegister() {

	wsocket.Register("login", controller.LoginController)

}
