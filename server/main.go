package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gilwong00/go-webrtc-app/pkg/room"
	"github.com/gilwong00/go-webrtc-app/pkg/routes"
	"github.com/gin-gonic/gin"
)

const (
	port = ":5000"
)

func main() {
	room.Rooms.Init()

	r := gin.Default()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	routes.AppRoutes(r)

	err := r.Run(port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(0)
	}
}
