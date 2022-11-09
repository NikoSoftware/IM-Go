package main

import (
	"IM-Go/routers"
	"IM-Go/viper"
	"IM-Go/wsocket"
)

func main() {
	viper.InitViper()

	wsocket.NewClientManagerInit()
	routers.WSocketHandlerRegister()
	go wsocket.DisposeChanService.Start()

	wsocket.StartWsocket()

}
