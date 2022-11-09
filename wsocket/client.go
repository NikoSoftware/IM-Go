package wsocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"runtime/debug"
)

type Client struct {
	Addr   string          //客户地址
	Conn   *websocket.Conn //连接
	UserId int
}

type Login struct {
	UserId int
	Client *Client
}

// 创建连接
func NewClient(addr string, conn *websocket.Conn) *Client {

	return &Client{
		Addr: addr,
		Conn: conn,
	}
}

func (c *Client) ReadMsg() {

	defer func() {

		if err := recover(); err != nil {
			fmt.Println("write stop", string(debug.Stack()), err)
		}
	}()

	defer func() {
		c.Conn.Close()
	}()

	for {

		_, p, err := c.Conn.ReadMessage()

		if err != nil {
			log.Printf("%s:消息读取异常", c.Addr)
		}

		MsgProcess(c, p)

	}

}

func (c *Client) WriteMsg() {

}

func (c *Client) SendMsg(msg []byte) {

}
