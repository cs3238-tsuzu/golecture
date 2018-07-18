package main

import (
	"strconv"
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

	write := func(msg string) {
		if err := conn.WriteJSON(Message{
			Message: msg,
			User:    "Prime Bot",
		}); err != nil {
			panic(err)
		}
	}

	var msg Message
MAINLOOP:
	for {
		if err := conn.ReadJSON(&msg); err != nil {
			panic(err)
		}

		if strings.HasPrefix(msg.Message, "/prime ") {
			str := strings.TrimPrefix(msg.Message, "/prime ")

			n, err := strconv.Atoi(str)

			if err != nil {
				write("Invalid format: " + err.Error() + "\nex. /prime 100")

				continue
			}

			if n <= 0 {
				write("Must be a positive number" + "\nex. /prime 100")

				continue
			}

			k := int64(n)
			if k != 2 && k%2 == 0 {
				write("Not Prime")
				continue
			}
			for i := int64(3); i*i < k; i += 2 {
				if k%i == 0 {
					write("Not Prime")
					continue MAINLOOP
				}
			}

			write("Prime")
		}
	}
}
