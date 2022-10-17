package room

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// create a new room and return the room id
func CreateRoomHandler(c *gin.Context) {}

// join a existing room
func JoinRoomHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "join",
	})
}
