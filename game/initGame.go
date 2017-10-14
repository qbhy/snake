package game

import (
	"github.com/gorilla/websocket"
)

func InitGame(ws *websocket.Conn, q interface{}) {
	name := q.(string)
	if Users != nil && Users[name] == name {
		ws.WriteJSON(Message{
			Action: "HandleError",
			Data:   "该名称已存在，请使用其他名称吧~",
		})
		return
	}

	Users[name] = name
	Clients[ws] = name
	ws.WriteJSON(Message{
		Action: "InitName",
		Data:   name,
	})
	PushMessage(Message{
		Action: "SetRoomInfo",
		Data:   SnakeRoom,
	})
}

func Ready(ws *websocket.Conn, q interface{}){
	name := Clients[ws]
	SnakeRoom.Snakes[name] = GenerateSnake(name)
	PushMessage(Message{
		Action: "SetRoomInfo",
		Data: SnakeRoom,
	})
}
