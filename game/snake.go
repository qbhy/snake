package game

func GenerateSnake(name string) Snake {
	l := len(SnakeRoom.Snakes)
	n := l * X * 2
	return Snake{
		Status:        READY,
		Name:          name,
		Speed:         1,
		PrevDirection: "right",
		NextDirection: "right",
		Body:          []int{6 + n, 5 + n, 4 + n, 3 + n, 2 + n, 1 + n},
		Color:         ColorMap[l],
	}
}
