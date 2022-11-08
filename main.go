package main

import (
	"IM-Go/viper"
	"IM-Go/wsocket"
)

func main() {

	viper.InitViper()
	wsocket.StartWsocket()

}
