package main

type Message struct {
	messageType int    `json:"message_type"`
	targetID    string `json:"target_id"`
	userID      string `json:"user_id"`
	roomID      string `json:"room_id"`
	Body        string `json:"body"`
	Timestamp   string `json:"timestamp"`
}
