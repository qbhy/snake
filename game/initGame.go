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
	UpdateRoom()
}

func Ready(ws *websocket.Conn, q interface{}) {
	name := Clients[ws]
	if len(SnakeRoom.Snakes) > 5 {
		SendError(ws, "该房间人数已满!")
		return
	}
	if _, ok := SnakeRoom.Snakes[name]; ok {
		SendError(ws, "您已准备好，请等待其他玩家加入!")
		return
	}
	SnakeRoom.Snakes[name] = GenerateSnake(name, len(SnakeRoom.Snakes))
	UpdateRoom()
}
