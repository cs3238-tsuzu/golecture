package main

import (
	"strings"

	"github.com/gorilla/websocket"
)

type Message struct {
	Message string `json:"message"`
	User    string `json:"user"`
}

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	//conn, _, err := websocket.DefaultDialer.Dial("wss://gochat.modoki.tsuzu.xyz/ws", nil)

	if err != nil {
		panic(err)
	}

	var msg Message
	for {
		if err := conn.ReadJSON(&msg); err != nil {
			panic(err)
		}

		if strings.HasPrefix(msg.Message, "/echo ") {
			msg.Message = strings.TrimPrefix(msg.Message, "/echo ")
			msg.User = "Echo Bot"

			if err := conn.WriteJSON(msg); err != nil {
				panic(err)
			}

		}

	}
}
