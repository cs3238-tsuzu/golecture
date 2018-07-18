package main

import (
	"log"
	"strconv"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message struct {
	Message string `json:"message"`
	User    string `json:"user"`
	UID     string `json:"uid"`
	ID      string `json:"id"`
}

type Hub struct {
	messageChan  chan Message
	registerChan chan *websocket.Conn
}

func (h *Hub) run() {
	var conns []*websocket.Conn

	idCounter := 1

	for {
		select {
		case msg := <-h.messageChan:
			removed := []int{}

			msg.ID = strconv.Itoa(idCounter)
			idCounter++

			for i := range conns {
				if err := conns[i].WriteJSON(msg); err != nil {
					removed = append(removed, i)

					log.Println("Unregistered")
				}
			}

			newConns := make([]*websocket.Conn, 0, len(conns)-len(removed))

			j := 0
			for i := range conns {
				if len(removed) > j && removed[j] == i {
					j++
				} else {
					newConns = append(newConns, conns[i])
				}
			}

		case r := <-h.registerChan:
			conns = append(conns, r)
		}
	}
}

func (h *Hub) register(c *websocket.Conn) {
	h.registerChan <- c
}

func (h *Hub) broadcast(msg Message) {
	h.messageChan <- msg
}

func main() {
	r := gin.Default()
	upgrader := websocket.Upgrader{}

	hub := Hub{
		messageChan:  make(chan Message, 1024),
		registerChan: make(chan *websocket.Conn, 1024),
	}

	go hub.run()

	var uidCounter int32

	r.GET("/ws", func(c *gin.Context) {
		log.Println("Connected")
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

		if err != nil {
			log.Println(err)
			return
		}

		uid := atomic.AddInt32(&uidCounter, 1)

		go func() {
			defer conn.Close()

			hub.register(conn)

			for {
				var msg Message
				if err := conn.ReadJSON(&msg); err != nil {
					return
				}

				msg.UID = strconv.Itoa(int(uid))

				hub.broadcast(msg)
			}
		}()
	})

	r.StaticFile("/", "./files/index.html")
	r.Static("/static", "./files/static")

	r.Run()
}
