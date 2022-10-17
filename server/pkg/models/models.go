package models

import "github.com/gorilla/websocket"

type Participant struct {
	Host bool
	Conn *websocket.Conn
}

type BroadcastMessage struct {
	Message map[string]interface{}
	RoomId  string
	Client  *websocket.Conn
}
