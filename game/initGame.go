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

func Entry(ws *websocket.Conn) {
	name := Clients[ws]
	if SnakeRoom.Status == RUNNING {
		SendError(ws, "游戏已开始，无法加入!")
		return
	}
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

func Ready(ws *websocket.Conn) {
	name := Clients[ws]
	if _, ok := SnakeRoom.Snakes[name]; !ok {
		SendError(ws, "您还没有加入游戏，请先加入游戏后准备！")
		return
	}
	snake := SnakeRoom.Snakes[name]
	if snake.Status == READY {
		SendError(ws, "您已经准备好，无需重复准备！")
		return
	}
	snake.Status = READY
	SnakeRoom.Snakes[name] = snake
	UpdateRoom()
	autoStart()
}

func autoStart() {
	for _, snake := range SnakeRoom.Snakes {
		if snake.Status != READY {
			return
		}
	}
	StartGame()
}

func StartGame() {
	SnakeRoom.Status = RUNNING
	PushMessage(Message{
		Action: "StartGame",
		Data:   SnakeRoom,
	})
}
