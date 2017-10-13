package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"snake/game"
	"fmt"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		println(r.Host)
		return true
	},
}

type WebSocketController struct {
	beego.Controller
}

func (this *WebSocketController) View() {
	this.TplName = "webSocket.tpl"
}

func (this *WebSocketController) WS() {

	ws, err := Upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	game.Clients[ws] = string(len(game.Clients))

	// 不断的从页面上获取数据
	for {
		var request game.Request
		err := ws.ReadJSON(&request)
		if err != nil {
			fmt.Println("Request")
			log.Printf("error: %v", err)
			delete(game.Clients, ws)
			break
		}
		game.HandleRequest(ws, request)
	}
}
