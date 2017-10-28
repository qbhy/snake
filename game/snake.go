package game

var ColorMap = []string{"red", "blue", "yellow", "green", "#ccc", "#399"}

func GenerateSnake(name string, len int) Snake {
	n := len * X * 3
	return Snake{
		Status:        WAITING,
		Name:          name,
		Speed:         1,
		PrevDirection: "right",
		NextDirection: "right",
		Body:          []int{6 + n, 5 + n, 4 + n, 3 + n, 2 + n, 1 + n},
		Color:         ColorMap[len],
	}
}
