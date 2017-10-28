package game

import (
	"github.com/gorilla/websocket"
	"log"
	"fmt"
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
	} else if request.Action == "Entry" {
		Entry(ws)
	} else if request.Action == "Ready" {
		Ready(ws)
	}
}

// 推送消息到 Broadcast
func PushMessage(message Message) {
	Broadcast <- message
}

// 用户关闭链接事件
func OnClose(ws *websocket.Conn) {
	name := Clients[ws]
	delete(Clients, ws)

	if _, ok := SnakeRoom.Snakes[name]; ok {
		snake := SnakeRoom.Snakes[name]
		delete(SnakeRoom.Snakes, name)
		if SnakeRoom.Status == WAITING {
			// 处理其他玩家的游戏状态
			i := 0
			for k := range SnakeRoom.Snakes {
				SnakeRoom.Snakes[k] = GenerateSnake(k, i)
				i++
			}
		} else if SnakeRoom.Status == RUNNING {
			// 正在游戏中断开链接的话让该用户的身体变成食物
			fmt.Println(snake)
		}
		PushMsg(name + "退出游戏")
		PushLog(name + "退出游戏")
	}

	delete(Users, name)
	UpdateRoom()
}

// 向客户端更新房间信息
func UpdateRoom() {
	PushMessage(Message{
		Action: "SetRoomInfo",
		Data:   SnakeRoom,
	})
}

// 广播发送至页面
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
