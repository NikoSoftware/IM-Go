package wsocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"net/http"
)

func StartWsocket() {

	http.HandleFunc("/ws", WsEndpoint)

	httpPort := viper.GetString("app.httpPort")

	http.ListenAndServe(":"+httpPort, nil)
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {

		fmt.Println(err)
	}

	client := NewClient(conn.RemoteAddr().String(), conn)

	go client.ReadMsg()

	go client.WriteMsg()

}
