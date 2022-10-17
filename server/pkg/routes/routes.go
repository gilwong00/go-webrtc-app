package routes

import (
	"github.com/gilwong00/go-webrtc-app/pkg/room"
	"github.com/gin-gonic/gin"
)

func AppRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/create", room.CreateRoomHandler)
	api.GET("/join", room.JoinRoomHandler)
}
