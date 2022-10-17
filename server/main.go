package main

import (
	"log"
	"os"

	"github.com/gilwong00/go-webrtc-app/pkg/room"
	"github.com/gilwong00/go-webrtc-app/pkg/routes"
	"github.com/gin-gonic/gin"
)

const (
	port = ":5000"
)

func main() {
	r := gin.Default()
	room.Rooms.Init()

	routes.AppRoutes(r)

	err := r.Run(port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(0)
	}
}
