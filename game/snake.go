package game

import (
	"github.com/gorilla/websocket"
	"log"
)

type Message struct {
	Message string `json:"message"`
	Data    string `json:"data"`
	Code    string `json:"code"`
}

type Request struct {
	Action string `json:"action"`
	Args   string `json:"args"`
}

var Clients = make(map[*websocket.Conn]bool)
var Broadcast = make(chan Message)

func init() {
	go HandleMessages()
}

func HandleRequest(request Request) {
	msg := Message{
		Message: "啦啦啦",
		Data:    request.Action,
	}
	Broadcast <- msg
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
