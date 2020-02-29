package ws_chat

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	sendTo chan message
}

func NewHub() *Hub {
	return &Hub{
		sendTo:     make(chan message),
		broadcast:  make(chan message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case message := <-h.sendTo:
			for client := range h.clients {
				if client.uuid == message.to {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}

func (h *Hub) Broadcast(str interface{}, event MessageEvent) {
	msg := message{
		Event: event,
		Data:  str,
	}

	h.broadcast <- msg
}

func (h *Hub) Send(uuid string, str interface{}, event MessageEvent) {
	msg := message{
		to:    uuid,
		Event: event,
		Data:  str,
	}

	h.sendTo <- msg
}

//Count get Clients
func (h *Hub) Count() int {
	return len(h.clients)
}
