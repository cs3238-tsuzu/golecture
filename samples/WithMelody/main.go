package main

import (
	"encoding/json"
	"log"
	"strconv"
	"sync/atomic"

	"gopkg.in/olahol/melody.v1"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
	User    string `json:"user"`
	UID     string `json:"uid"`
	ID      string `json:"id"`
}

func main() {
	r := gin.Default()
	m := melody.New()
	m.Config.MaxMessageSize = 10 * 1024

	var uidCounter, idCounter int32

	r.GET("/ws", func(c *gin.Context) {
		log.Println("Connected")

		uid := atomic.AddInt32(&uidCounter, 1)
		m.HandleRequestWithKeys(c.Writer, c.Request, map[string]interface{}{"uid": int(uid)})
	})

	m.HandleMessage(func(s *melody.Session, bytes []byte) {
		uid := strconv.Itoa(s.Keys["uid"].(int))

		var msg Message
		if err := json.Unmarshal(bytes, &msg); err != nil {
			return
		}

		msg.UID = uid
		msg.ID = strconv.Itoa(int(atomic.AddInt32(&idCounter, 1)))

		b, _ := json.Marshal(msg)
		m.Broadcast(b)
	})

	r.StaticFile("/", "./files/index.html")
	r.Static("/static", "./files/static")

	r.Run()
}
