package models

import "github.com/gorilla/websocket"

type Participant struct {
	Host bool
	Conn *websocket.Conn
}
