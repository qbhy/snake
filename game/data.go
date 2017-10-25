package game

import (
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
)

var Redis *redis.Client

const (
	X       = 60
	Y       = 30
	SPEED   = 1000
	WAITING = "waiting"
	RUNNING = "running"
	READY   = "ready"
)

// 蛇
type Snake struct {
	Status        string `json:"status"`
	Name          string `json:"name"`
	Speed         int    `json:"speed"`
	PrevDirection string `json:"prev_direction"`
	NextDirection string `json:"next_direction"`
	Body          []int  `json:"body"`
	Color         string `json:"color"`
}

// 房间
type Room struct {
	Status     string           `json:"status"`
	X          int              `json:"x"`
	Y          int              `json:"y"`
	Speed      int              `json:"speed"`
	Snakes     map[string]Snake `json:"snakes"`
	Rule       Rule             `json:"rule"`
	Spectators []string         `json:"spectators"`
	Foods      []int            `json:"foods"`
	//Logs       []string         `json:"logs"`
}

type Rule struct {
	Top    int `json:"top"`
	Right  int `json:"right"`
	Bottom int `json:"bottom"`
	Left   int `json:"left"`
}

var Users = map[string]string{}

var SnakeRoom Room

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	SnakeRoom = Room{
		Status: WAITING,
		X:      X,
		Y:      Y,
		Speed:  SPEED,
		Rule: Rule{
			Top:    -X,
			Right:  1,
			Bottom: X,
			Left:   -1,
		},
		Snakes:     make(map[string]Snake),
		Spectators: []string{},
		Foods:      []int{},
	}

}

func AddLog(ws *websocket.Conn, log interface{}) {
	PushLog(Clients[ws] + "说:" + log.(string))
}
