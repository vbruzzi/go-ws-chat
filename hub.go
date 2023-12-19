package main

type Hub struct {
	clients   map[*Client]bool
	register  chan *Client
	broadcast chan []byte
}

func (h *Hub) init() {
	for {
		select {
		case newClient := <-h.register:
			h.clients[newClient] = true
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.inbound <- message:
				}
			}
		}
	}
}

func newHub() *Hub {
	return &Hub{
		make(map[*Client]bool),
		make(chan *Client),
		make(chan []byte),
	}
}
