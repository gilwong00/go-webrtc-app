package main

import (
	"github.com/gilwong00/go-webrtc-app/pkg/routes"
	"github.com/gin-gonic/gin"
)

const (
	port = ":5000"
)

func main() {
	r := gin.Default()

	routes.AppRoutes(r)

	r.Run(port)
}
