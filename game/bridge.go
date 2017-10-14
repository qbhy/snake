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
	Args   interface{} `json:"args"`
}

var Status = "waiting"

var Clients = make(map[*websocket.Conn]string)
var Broadcast = make(chan Message)

func init() {
	go HandleMessages()
}

func HandleRequest(ws *websocket.Conn, request Request) {
	if request.Action == "init" {
		// init
	} else if request.Action == "InitGame" {
		InitGame(ws, request.Args)
	} else if request.Action == "AddLog" {
		AddLog(ws, request.Args)
	} else if request.Action == "Ready" {
		Ready(ws, request.Args)
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
