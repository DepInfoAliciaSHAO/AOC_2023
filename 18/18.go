package main

import (
	"AOC2023/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func getVertices(input []string) ([]utils.Vertex, int) {
	var vertices = make([]utils.Vertex, 0)
	var current = utils.Vertex{}
	var currentLength = 0
	vertices = append(vertices, current)
	for _, line := range input {
		var n, _ = strconv.Atoi(strings.TrimSpace(line[2:4]))
		currentLength += n
		switch line[0:1] {
		case "R":
			current = current.Move(0, n)
		case "L":
			current = current.Move(0, -n)
		case "U":
			current = current.Move(-n, 0)
		case "D":
			current = current.Move(n, 0)
		default:
			panic("Error reading direction")
		}
		vertices = append(vertices, current)
	}
	return vertices[0 : len(vertices)-1], currentLength
}

func partOne(input string) int {
	var vertices, length = getVertices(utils.LineByLine(input))
	return utils.Pick(utils.Shoelace(vertices), length)
}

func getVertices2(input []string) ([]utils.Vertex, int) {
	var vertices = make([]utils.Vertex, 0)
	var current = utils.Vertex{}
	var zero64 = 0
	var currentLength = zero64
	vertices = append(vertices, current)
	for _, line := range input {
		var components = strings.Split(line, " ")
		var hexadecimal = components[2][2 : len(components[2])-1]
		var m, _ = strconv.ParseInt(hexadecimal[0:len(hexadecimal)-1], 16, 64)
		var direction64, _ = strconv.ParseInt(hexadecimal[len(hexadecimal)-1:], 16, 64)
		var direction = int(direction64)
		var n = int(m)
		currentLength += n
		switch direction {
		case 0:
			current = current.Move(zero64, n)
		case 2:
			current = current.Move(zero64, -n)
		case 3:
			current = current.Move(-n, zero64)
		case 1:
			current = current.Move(n, zero64)
		default:
			panic("Error reading direction")
		}
		vertices = append(vertices, current)
	}
	return vertices[0 : len(vertices)-1], currentLength
}

func partTwo(input string) int {
	var vertices, length = getVertices2(utils.LineByLine(input))
	return utils.Pick(utils.Shoelace(vertices), length)
}

func main() {
	var start = time.Now()
	fmt.Println(partOne("18/input.txt"))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(partTwo("18/input.txt"))
	fmt.Println(time.Since(start))
}
