package main

import (
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	Message string `json:"message"`
	User    string `json:"user"`
}

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)

	if err != nil {
		panic(err)
	}

	var msg Message
	for {
		if err := conn.ReadJSON(&msg); err != nil {
			panic(err)
		}

		if msg.Message == "/time" {
			now := time.Now()
			msg.Message = now.Format(time.RFC3339)
			msg.User = "Echo Bot"

			if err := conn.WriteJSON(msg); err != nil {
				panic(err)
			}

		}

	}
}
