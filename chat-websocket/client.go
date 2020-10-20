package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	userID string
	// socket is the web socket for this client.
	socket *websocket.Conn
	// forward is a channel that holds incoming messages
	// forward chan *Message
	// send is a channel on which messages are sent.
	send chan *Message
	// register is a channel for client wishing to join the specific client.
	// register chan *Connection
	// unregister is a channel for client wishing to rmeove the specific client.
	// unregister chan *Connection
	// holds all current clients in this room
	// clientID uuid.UUID
	// Registered clients.
	// clients map[uuid.UUID]map[*Connection]bool
	// room is the room this client is chatting in.
	rooms map[string]*Room
	// room is the specific room this client want to join in.
	roomID string
}

func newClient(ws *websocket.Conn, userID string) *Client {
	return &Client{
		// forward:    make(chan *Message),
		// register:   make(chan *Connection),
		// unregister: make(chan *Connection),
		// clients:    make(map[uuid.UUID]map[*Connection]bool),
		userID: userID,
		socket: ws,
		send:   make(chan *Message),
		rooms:  make(map[string]*Room),
	}
}

// readPump pumps messages from the websocket connection to the hub.
func (c *Client) readPump() {
	defer func() {
		c.socket.Close()
	}()

	c.socket.SetReadLimit(maxMessageSize)
	c.socket.SetReadDeadline(time.Now().Add(pongWait))
	c.socket.SetPongHandler(func(string) error {
		c.socket.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		m := Message{}
		err := c.socket.ReadJSON(&m)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}

			break
		}

		c.rooms[c.roomID].forward <- &m
	}
}

// writePump pumps messages from the hub to the websocket connection.
func (c *Client) writePump() {
	m := Message{}
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, &m)
				return
			}

			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, &m); err != nil {
				return
			}
		}
	}
}

// write writes a message with the given message type and payload.
func (c *Client) write(mt int, payload *Message) error {
	c.socket.SetWriteDeadline(time.Now().Add(writeWait))
	payload.messageType = mt
	return c.socket.WriteJSON(payload)
}
