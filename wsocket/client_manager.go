package wsocket

type ClientManager struct {
	Clients map[string]*Client
	Users   map[int]*Client
}

type DisposeChan struct {
	Register chan *Client

	Login chan *login

	Unregister chan *Client
}
