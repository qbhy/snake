package game

import (
	"github.com/gorilla/websocket"
)

// 蛇
type Snake struct {
	Status        string `json:"status"`
	Name          string `json:"name"`
	Speed         int    `json:"speed"`
	PrevDirection string `json:"prev_direction"`
	NextDirection string `json:"next_direction"`
	Body          []int  `json:"body"`
	Color         []int  `json:"color"`
}


type State struct {
	Status     string   `json:"status"`
	X          int      `json:"x"`
	Y          int      `json:"y"`
	Speed      int      `json:"speed"`
	Snakes     []Snake  `json:"snakes"`
	Rule       Rule     `json:"rule"`
	Spectators []string `json:"spectators"`
	Foods      []int    `json:"foods"`
	Logs       []string `json:"logs"`
}

func initGame(ws *websocket.Conn) {
	ws.WriteJSON(Message{
		Action: "setInitState",
		Data:   SnakeRoom,
	})
}

func initName(ws *websocket.Conn, q interface{}) {
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
		Action: "SetName",
		Data:   name,
	})
}
