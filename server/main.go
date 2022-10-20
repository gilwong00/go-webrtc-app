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
	port = ":8000"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	room.Rooms.Init()

	r := gin.Default()
	r.Use(CORSMiddleware())
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
