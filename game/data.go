package game

import (
	"github.com/go-redis/redis"
)

var Redis *redis.Client

const (
	X     = 30
	Y     = 30
	SPEED = 1000
)

type Room struct {
	Status     string           `json:"status"`
	X          int              `json:"x"`
	Y          int              `json:"y"`
	Speed      int              `json:"speed"`
	Snakes     map[string]Snake `json:"snakes"`
	Rule       Rule             `json:"rule"`
	Spectators []string         `json:"spectators"`
	Foods      []int            `json:"foods"`
	Logs       []string         `json:"logs"`
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
		Status: "waiting",
		X:      X,
		Y:      Y,
		Speed:  SPEED,
		Rule: Rule{
			Top:    -X,
			Right:  1,
			Bottom: X,
			Left:   -1,
		},
		Snakes: map[string]Snake{},
		Logs:   []string{},
		Foods:  []int{},
	}

}
