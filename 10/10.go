package main

import (
	"AOC2023/utils"
	"bufio"
	"fmt"
	"os"
	"time"
)

/*
Returns a 2D grid map, mapping the field map, the dimensions of the grid and the position of the starting point.
*/
func getPipes(input string) (utils.Grid[string], int, int, utils.Vertex) {
	// Open the file
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, 0, 0, utils.Vertex{}
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)
	var grid = make(utils.Grid[string])
	var i = 0
	var k = 0
	var start = utils.Vertex{}
	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			k = len(line)
		}
		for j, _ := range line {
			var pipe = line[j : j+1]
			grid[utils.Vertex{X: i, Y: j}] = pipe
			if pipe == "S" {
				start.X = i
				start.Y = j
			}
		}
		i++
	}
	return grid, i, k, start
}

type Field struct {
	startingPoint utils.Vertex
	startingPipe  string
	grid          utils.Grid[string]
}

/*
Returns a list of the neighbors connected to a pipe.
*/
func validNeighbors(vertex utils.Vertex, field Field) []utils.Vertex {
	switch field.grid[vertex] {
	case "|":
		return []utils.Vertex{vertex.UpNeighbor(), vertex.DownNeighbor()}
	case "-":
		return []utils.Vertex{vertex.RightNeighbor(), vertex.LeftNeighbor()}
	case "L":
		return []utils.Vertex{vertex.UpNeighbor(), vertex.RightNeighbor()}
	case "J":
		return []utils.Vertex{vertex.UpNeighbor(), vertex.LeftNeighbor()}
	case "7":
		return []utils.Vertex{vertex.LeftNeighbor(), vertex.DownNeighbor()}
	case "F":
		return []utils.Vertex{vertex.RightNeighbor(), vertex.DownNeighbor()}
	default:
		return nil
	}
}

/*
During a traversal of the loop, returns the coordinates of the next vertex to be explored.
*/
func nextVertex(currentVertex utils.Vertex, previousVertex utils.Vertex, field Field) utils.Vertex {
	var validNeighbors = validNeighbors(currentVertex, field)
	if validNeighbors[0] == previousVertex {
		return validNeighbors[1]
	} else {
		return validNeighbors[0]
	}
}

/*
Returns the two points next to the starting point. They will define the direction of traversal.
Also determines which pipe is the starting point and returns its string representation.
*/
func findStartingDirection(field Field) ([]utils.Vertex, string) {
	var directions = make([]utils.Vertex, 0)
	var directionStrings = "|7FLJ-"
	var directionCount = make([]int, len(directionStrings))
	var up, okUp = field.grid[field.startingPoint.UpNeighbor()]
	var left, okLeft = field.grid[field.startingPoint.LeftNeighbor()]
	var right, okRight = field.grid[field.startingPoint.RightNeighbor()]
	var down, okDown = field.grid[field.startingPoint.DownNeighbor()]
	if okUp && (up == "|" || up == "7" || up == "F") {
		directions = append(directions, field.startingPoint.UpNeighbor())
		directionCount[0] += 1
		directionCount[1] += 3
		directionCount[2] += 4
	}
	if okLeft && (left == "-" || left == "L" || left == "F") {
		directions = append(directions, field.startingPoint.LeftNeighbor())
		directionCount[0] += 2
		directionCount[1] += 5
		directionCount[2] += 4
	}
	if okRight && (right == "-" || right == "J" || right == "7") {
		directions = append(directions, field.startingPoint.RightNeighbor())
		directionCount[5] += 1
		directionCount[2] += 1
		directionCount[3] += 1
	}
	if okDown && (down == "|" || down == "L" || down == "J") {
		directions = append(directions, field.startingPoint.DownNeighbor())
		directionCount[0] += 1
		directionCount[1] += 1
		directionCount[2] += 1
	}
	var index = 0
	for i, v := range directionCount {
		if v == 2 {
			index = i
		}
	}
	return directions, directionStrings[index : index+1]
}

/*
Returns the farthest distance from the starting point of a loop in a given field.
The distance is computed by doing the two possible traversal of the loop at the same time.
The point at which the farthest distance = k is reached is at the unique point where after k steps,
both traversals are at the same point.
*/
func loop(field Field) int {
	var startingDirections, _ = findStartingDirection(field)
	var previousA, previousB = field.startingPoint, field.startingPoint
	var A = startingDirections[0]
	var B = startingDirections[1]
	var distance = 1
	for A != B {
		var currentA = A
		var currentB = B
		A = nextVertex(A, previousA, field)
		B = nextVertex(B, previousB, field)
		previousA = currentA
		previousB = currentB
		distance += 1
	}
	return distance
}

func partOne(input string) int {
	var pipes, _, _, startPoint = getPipes(input)
	var start = startPoint
	return loop(Field{startingPoint: start, grid: pipes})
}

type Loop struct {
	field Field
	loop  utils.Grid[string]
	dimX  int
	dimY  int
}

/*
From a field where there is a loop, returns a 2D grid map of the coordinates of the points that make up the loop.
*/
func getLoop(field Field) utils.Grid[string] {
	var loop = make(utils.Grid[string])
	loop[field.startingPoint] = field.grid[field.startingPoint]
	var startingDirections, _ = findStartingDirection(field)
	var previousA = field.startingPoint
	var A = startingDirections[0]
	for A != field.startingPoint {
		loop[A] = field.grid[A]
		var currentA = A
		A = nextVertex(A, previousA, field)
		previousA = currentA
	}
	return loop
}

/*
Finds the number of points inside a loop on a given line of the grid.
A ray is cast from left to right. It's inside or outside the loop, and this state changes
when the ray encounters a pipe from the loop.
*/
func ray(lineIndex int, loop Loop) int {
	var in = false
	var nInside = 0

	for j := 0; j < loop.dimY; j++ {
		var point = utils.Vertex{X: lineIndex, Y: j}
		var direction, onLoop = loop.loop[point]
		// If S is reached, it needs to be replaced by its pipe.
		if loop.loop[point] == "S" {
			direction = loop.field.startingPipe
		}
		// Only up or (exclusive or) down pipes change the ray's state.
		if onLoop {
			if direction == "|" || direction == "L" || direction == "J" {
				in = !in
			}
		} else {
			if in && !onLoop {
				nInside += 1
			}
		}
	}
	return nInside
}

func partTwo(input string) int {
	var pipes, dimX, dimY, startPoint = getPipes(input)
	var start = startPoint
	var field = Field{startingPoint: start, grid: pipes}
	var loop = Loop{field, getLoop(field), dimX, dimY}
	var _, pipe = findStartingDirection(loop.field)
	loop.field.startingPipe = pipe
	var res = 0
	for i := 0; i < dimX; i++ {
		res += ray(i, loop)
	}
	return res
}
func main() {
	var start = time.Now()
	fmt.Println(partOne("10/input.txt"))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(partTwo("10/input.txt"))
	fmt.Println(time.Since(start))
}
