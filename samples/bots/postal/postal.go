package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/japanese"

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
			User:    "Postal Bot",
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

		if strings.HasPrefix(msg.Message, "/postal ") {
			postal := strings.TrimPrefix(msg.Message, "/postal ")

			if len(postal) > 7 || len(postal) < 3 {
				write("Invalid format")
				continue
			}

			for i := range postal {
				if postal[i] < '0' || postal[i] > '9' {
					write("Invalid format")
					continue MAINLOOP
				}
			}

			resp, err := http.DefaultClient.Get("http://zip.cgis.biz/csv/zip.php?zn=" + postal)

			if err != nil {
				write("API error: " + err.Error())
				continue
			}

			b, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()

			if err != nil {
				write("API error: " + err.Error())
				continue
			}

			var str = string(b)
			str = strings.Join(strings.Split(str, "\r\n"), ",")
			str = strings.Join(strings.Split(str, "\n"), ",")
			str = strings.Join(strings.Split(str, "\r"), ",")

			arr := strings.Split(str, ",")

			for i := range arr {
				if arr[i] == `"none"` {
					arr[i] = ""
				}

				arr[i] = strings.TrimPrefix(arr[i], `"`)
				arr[i] = strings.TrimSuffix(arr[i], `"`)
				arr[i], _ = japanese.EUCJP.NewDecoder().String(arr[i])
			}

			if arr[8] == "0" {
				write("Not found")

				continue
			}

			msg := strings.Join(arr[len(arr)-5:], "")

			write(msg)
		}
	}
}
