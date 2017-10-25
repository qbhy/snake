package game

import "github.com/gorilla/websocket"

func PushMsg(msg string) {
	PushMessage(Message{
		Action: "ShowMessage",
		Data:   msg,
	})
}

func PushLog(log string) {
	PushMessage(Message{
		Action: "AddLog",
		Data:   log,
	})
}

func SendError(ws *websocket.Conn, err string) {
	ws.WriteJSON(Message{
		Action: "HandleError",
		Data:   err,
	})
}