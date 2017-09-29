package game

import (
	"github.com/gorilla/websocket"
	"fmt"
)

const (
	X     = 30
	Y     = 30
	SPEED = 1000
)

// 蛇
type Snake struct {
	Status        string `json:"status"`
	Name          string `json:"name"`
	Speed         int    `json:"speed"`
	PrevDirection string `json:"prev_direction"`
	NextDirection string `json:"next_direction"`
	Body          []int  `json:"body"`
}

type Rule struct {
	Top    int `json:"top"`
	Right  int `json:"right"`
	Bottom int `json:"bottom"`
	Left   int `json:"left"`
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

	state := State{
		X:          X,
		Y:          Y,
		Status:     Status,
		Speed:      SPEED,
		Foods:      []int{},
		Spectators: []string{},
		Logs:       []string{},
		Rule: Rule{
			Top:    -X,
			Right:  1,
			Bottom: X,
			Left:   -1,
		},
		Snakes: []Snake{
			Snake{
				Name: "qbhy",
				Speed: 1,
			},
		},
	}
	fmt.Println(state)

	ws.WriteJSON(state)

}
