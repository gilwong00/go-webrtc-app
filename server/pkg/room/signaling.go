package room

import (
	"log"
	"net/http"

	"github.com/gilwong00/go-webrtc-app/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var Rooms RoomMap
var broadcastChannel = make(chan models.BroadcastMessage)

// create a new room and return the room id
func CreateRoomHandler(c *gin.Context) {
	roomId := Rooms.CreateRoom()

	c.JSON(http.StatusOK, gin.H{
		"roomId": roomId,
	})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func broadcast() {
	for {
		msg := <-broadcastChannel

		for _, client := range Rooms.Map[msg.RoomId] {
			if client.Conn != msg.Client {
				err := client.Conn.WriteJSON(msg.Message)
				if err != nil {
					log.Fatalf("Failed to write to connection: %v", err)
					client.Conn.Close()
				}
			}
		}
	}
}

// join a existing room
func JoinRoomHandler(c *gin.Context) {
	roomId := c.Query("roomId")

	if len(roomId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Room Id cannot be empty",
		})
		return
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Fatalf("Websocket upgrade failed: %v", err)
	}

	Rooms.AddParticipantToRoom(roomId, false, ws)

	go broadcast()

	for {
		var msg models.BroadcastMessage

		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Fatalf("Read err: %v", err)
		}
		msg.Client = ws
		msg.RoomId = roomId

		broadcastChannel <- msg
	}
}
