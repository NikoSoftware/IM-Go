package wsocket

import "log"

var (
	ClientManagerService *ClientManager
	DisposeChanService   *DisposeChan
)

type ClientManager struct {
	Clients map[string]*Client
	Users   map[int]*Client
}

type DisposeChan struct {
	Register chan *Client

	Login chan *Login

	Unregister chan *Client
}

// 初始化客户端管理
func NewClientManagerInit() {

	ClientManagerService = &ClientManager{
		Clients: make(map[string]*Client, 1000),
		Users:   make(map[int]*Client, 1000),
	}

	DisposeChanService = &DisposeChan{

		Register:   make(chan *Client, 1000),
		Login:      make(chan *Login, 1000),
		Unregister: make(chan *Client, 1000),
	}

}

// 管道处理程序
func (manager *DisposeChan) Start() {
	for {
		select {
		case conn := <-manager.Register:
			// 建立连接事件
			manager.EventRegister(conn)

		case login := <-manager.Login:
			// 用户登录
			manager.EventLogin(login)

		case conn := <-manager.Unregister:
			// 断开连接事件
			manager.EventUnregister(conn)

			//case message := <-manager.Broadcast:
			//	// 广播事件
			//	clients := manager.GetClients()
			//	for conn := range clients {
			//		select {
			//		case conn.Send <- message:
			//		default:
			//			close(conn.Send)
			//		}
			//	}
		}
	}
}

func (manager *DisposeChan) EventRegister(conn *Client) {

}

func (manager *DisposeChan) EventLogin(login *Login) {

	login.Client.IsLogin = true

	ClientManagerService.Users[login.UserId] = login.Client

	log.Printf("用户登陆：%v", login.UserId)

}

func (manager *DisposeChan) EventUnregister(conn *Client) {

}
