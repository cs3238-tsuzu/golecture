package main

import (
	"log"

	"gopkg.in/olahol/melody.v1"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	m := melody.New()
	m.Config.MaxMessageSize = 10 * 1024

	r.GET("/ws", func(c *gin.Context) {
		log.Println("Connected")

		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	r.StaticFile("/", "./files/index.html")
	r.Static("/static", "./files/static")

	r.Run()
}
