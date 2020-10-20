package main

type Room struct {
	roomID string
	// forward is a channel that holds incoming messages
	forward chan *Message
	// join is a channel for clients wishing to join the room.
	join chan *Client
	// leave is a channel for clients wishing to leave the room
	leave chan *Client
	// holds all current clients in this room
	clients map[*Client]bool
}

func newRoom() *Room {
	return &Room{
		forward: make(chan *Message),
		join:    make(chan *Client),
		leave:   make(chan *Client),
		clients: make(map[*Client]bool),
	}
}

func (r *Room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward messsage to all clients
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}
