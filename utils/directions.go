package utils

// North: 0, East: 1, South: 2, West: 3

var DIRECTIONS = map[string]int{
	"North": 0,
	"East":  1,
	"South": 2,
	"West":  3,
}

var INTDIRECTIONS = map[int]string{
	0: "North",
	1: "East",
	2: "South",
	3: "West",
}

func ClockwiseRotation(direction string) string {
	return INTDIRECTIONS[(DIRECTIONS[direction]+1)%4]
}

func trigonometricRotation(direction string) string {
	return INTDIRECTIONS[(DIRECTIONS[direction]+4-1)%4]
}

func oppositeDirection(direction string) string {
	return INTDIRECTIONS[(DIRECTIONS[direction]-4-2)%4]
}
