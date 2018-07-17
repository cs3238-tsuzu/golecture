package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	Message string `json:"message"`
	User    string `json:"user"`
}

type FortuneTellingResult struct {
	Horoscope map[string][]struct {
		Content string `json:"content"`
		Item    string `json:"item"`
		Money   int    `json:"money"`
		Total   int    `json:"total"`
		Job     int    `json:"job"`
		Color   string `json:"color"`
		Love    int    `json:"love"`
		Rank    int    `json:"rank"`
		Sign    string `json:"sign"`
	} `json:"horoscope"`
}

func main() {
	//conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	conn, _, err := websocket.DefaultDialer.Dial("wss://gochat.modoki.tsuzu.xyz/ws", nil)

	if err != nil {
		panic(err)
	}
	write := func(msg string) {
		if err := conn.WriteJSON(Message{
			Message: msg,
			User:    "Fortune Telling Bot",
		}); err != nil {
			panic(err)
		}
	}

	var msg Message
	for {
		if err := conn.ReadJSON(&msg); err != nil {
			panic(err)
		}

		if msg.Message == "/fortune" {
			url := time.Now().Format("http://api.jugemkey.jp/api/horoscope/free/2006/01/02")

			resp, err := http.DefaultClient.Get(url)

			if err != nil {
				write("API error: " + err.Error())

				continue
			}

			var res FortuneTellingResult
			err = json.NewDecoder(resp.Body).Decode(&res)
			resp.Body.Close()

			if err != nil {
				write("API error: " + err.Error())

				continue
			}

			for _, v := range res.Horoscope {
				b, _ := json.MarshalIndent(v, "", "  ")

				write("占いの結果は・・・！\n" + string(b))
				break
			}

		}

	}
}
