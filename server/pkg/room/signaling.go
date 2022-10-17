package room

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Rooms RoomMap

// create a new room and return the room id
func CreateRoomHandler(c *gin.Context) {
	roomId := Rooms.CreateRoom()

	c.JSON(http.StatusOK, gin.H{
		"roomId": roomId,
	})
}

// join a existing room
func JoinRoomHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "join ping",
	})
}
