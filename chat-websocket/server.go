package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type ServeWs struct {
	addUser        chan *Client
	removeUser     chan *Client
	connectedUsers map[string]*Client
	addRoom        chan *Room
	removeRoom     chan *Room
	connectedRooms map[string]*Room
	assignToRoom   chan *Client
	unAssignToRoom chan *Client
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func newServeWs() *ServeWs {
	return &ServeWs{
		connectedUsers: make(map[string]*Client),
		addUser:        make(chan *Client),
		removeUser:     make(chan *Client),
		addRoom:        make(chan *Room),
		connectedRooms: make(map[string]*Room),
		removeRoom:     make(chan *Room),
		assignToRoom:   make(chan *Client),
		unAssignToRoom: make(chan *Client),
	}
}

// AddUser assign user to addUser channel
func (s *ServeWs) AddUser(c *Client) {
	s.addUser <- c
}

// RemoveUser assign user to connectedUsers channel
func (s *ServeWs) RemoveUser(c *Client) {
	s.removeUser <- c
}

// AddRoom create new room to connectedUsers channel
func (s *ServeWs) AddRoom(r *Room) {
	s.addRoom <- r
}

// RemoveRoom assign room is removed to connectedRooms channel
func (s *ServeWs) RemoveRoom(r *Room) {
	s.removeRoom <- r
}

// AssignToRoom assign user to connectedRooms channel
func (s *ServeWs) AssignToRoom(c *Client) {
	s.assignToRoom <- c
}

// UnAssignToRoom un-assign user from connectedRooms channel
func (s *ServeWs) UnAssignToRoom(c *Client) {
	s.unAssignToRoom <- c
}

// run assign user to channel
func (s *ServeWs) run() {
	for {
		select {
		case user := <-s.addUser:
			log.Println("Added a new User")
			s.connectedUsers[user.userID] = user
		case user := <-s.removeUser:
			log.Println("Remove a new User")
			delete(s.connectedUsers, user.userID)
		case room := <-s.addRoom:
			log.Println("Add a new Room")
			s.connectedRooms[room.roomID] = room
		case room := <-s.removeRoom:
			log.Println("Remove a new Room")
			delete(s.connectedRooms, room.roomID)
		case user := <-s.assignToRoom:
			if s.connectedRooms[user.roomID] == nil {
				return
			}

			s.connectedRooms[user.roomID].join <- user
		case user := <-s.unAssignToRoom:
			if s.connectedRooms[user.roomID] == nil {
				return
			}

			s.connectedRooms[user.roomID].leave <- user
		}
	}
}

// serveWsForAddUser handles add user to websocket.
func (s *ServeWs) serveWsForAddUser(w http.ResponseWriter, r *http.Request, userID string) {
	ws := originWs(w, r)

	log.Println("going to add user", userID)
	user := newClient(ws, userID)
	s.AddUser(user)
	log.Println("user added successfully")
}

// serveWsForRemoveUser
func (s *ServeWs) serveWsForRemoveUser(w http.ResponseWriter, r *http.Request, userID string) {
	ws := originWs(w, r)

	log.Println("going to remove user", userID)
	user := newClient(ws, userID)
	s.RemoveUser(user)
	log.Println("user removed successfully")
}

// serveWsForAddRoom
func (s *ServeWs) serveWsForAddRoom(w http.ResponseWriter, r *http.Request, roomID string) {
	log.Println("going to add room", roomID)
	room := newRoom()
	room.roomID = roomID
	s.AddRoom(room)
	log.Println("room added successfully")
}

// serveWsForAddRoom
func (s *ServeWs) serveWsForRemoveRoom(w http.ResponseWriter, r *http.Request, roomID string) {
	log.Println("going to remove room", roomID)
	room := newRoom()
	room.roomID = roomID
	s.RemoveRoom(room)
	log.Println("room removed successfully")
}

// serveWsForAddRoom
func (s *ServeWs) serveWsForAssignUserToRoom(w http.ResponseWriter, r *http.Request, userID, roomID string) {
	ws := originWs(w, r)

	log.Println("going to assign user to specific room", roomID)
	user := s.connectedUsers[userID]
	if user == nil {
		return
	}

	user = newClient(ws, userID)
	user.roomID = roomID
	s.AssignToRoom(user)
	log.Println("user assigned successfully")
	go user.writePump()
	go user.readPump()
}

// serveWsForAddRoom
func (s *ServeWs) serveWsForUnAssignUserFromRoom(w http.ResponseWriter, r *http.Request, userID, roomID string) {
	ws := originWs(w, r)

	log.Println("going to un-assign user from specific room", roomID)
	user := s.connectedUsers[userID]
	if user == nil {
		return
	}

	user = newClient(ws, userID)
	user.roomID = roomID
	s.UnAssignToRoom(user)
	log.Println("user un-assigned successfully")
}

func originWs(w http.ResponseWriter, r *http.Request) *websocket.Conn {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return ws
}
