package game

import (
	"github.com/gorilla/websocket"
	"log"
)

type Message struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Action  string      `json:"action"`
}

type Request struct {
	Action string      `json:"action"`
	Args  interface{} `json:"args"`
}

var Status = "waiting"

var Clients = make(map[*websocket.Conn]bool)
var Broadcast = make(chan Message)

func init() {
	go HandleMessages()
}

func HandleRequest(ws *websocket.Conn, request Request) {
	msg := Message{
		Message: "啦啦啦",
		Action:  request.Action,
	}
	if request.Action == "init" {
		initGame(ws)
	} else {
		PushMessage(msg)
	}
}

func PushMessage(message Message) {
	Broadcast <- message
}

//广播发送至页面
func HandleMessages() {
	for {
		msg := <-Broadcast
		for client := range Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("client.WriteJSON error: %v", err)
				client.Close()
				delete(Clients, client)
			}
		}
	}
}
