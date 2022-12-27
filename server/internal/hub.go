package internal

type Hub struct {
	clients    map[*Client]bool // 上线clients
	broadcast  chan string      // 客户端发送的消息 ->广播给其他的客户端
	register   chan *Client     // 注册channel，接收注册msg
	unregister chan *Client     // 下线channel
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan string),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		//用户加入房间
		case client := <-hub.register:
			hub.clients[client] = true
		//用户下线
		case client := <-hub.unregister:
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				close(client.send)

			}
		//广播消息
		case message := <-hub.broadcast:
			for client := range hub.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(hub.clients, client)
				}
			}
		}
	}
}
