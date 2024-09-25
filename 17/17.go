package main

import (
	"AOC2023/utils"
	"fmt"
	"strconv"
	"time"
)

type Node struct {
	position       utils.Vertex
	direction      int
	directionCount int
}
type Tile struct {
	node     Node
	priority int
}

// North: 0, East: 1, South: 2, West: 3
func clockwiseRotation(direction int) int {
	return (direction + 1) % 4
}

func trigonometricRotation(direction int) int {
	return (direction + 4 - 1) % 4
}

func oppositeDirection(direction int) int {
	return (direction - 4 - 2) % 4
}

func lessThan(t1 Tile, t2 Tile) bool {
	return t1.priority < t2.priority
}

func neighborFromDirection(v utils.Vertex, direction int) utils.Vertex {
	switch direction {
	case 0:
		return v.UpNeighbor()
	case 1:
		return v.RightNeighbor()
	case 2:
		return v.DownNeighbor()
	case 3:
		return v.LeftNeighbor()
	default:
		panic("Invalid direction.")
	}
}

func (t Tile) possibleNeighbors(graph utils.Grid[int], minForward int, maxForward int) []Node {
	var neighbors = make([]Node, 0)
	var cwr, tr = clockwiseRotation(t.node.direction), trigonometricRotation(t.node.direction)
	var n1 = neighborFromDirection(t.node.position, cwr)
	var _, ok1 = graph[n1]
	if ok1 && t.node.directionCount > minForward-1 {
		neighbors = append(neighbors, Node{n1, cwr, 1})
	}
	var n2 = neighborFromDirection(t.node.position, tr)
	var _, ok2 = graph[n2]
	if ok2 && t.node.directionCount > minForward-1 {
		neighbors = append(neighbors, Node{n2, tr, 1})
	}
	var n3 = neighborFromDirection(t.node.position, t.node.direction)
	var _, ok3 = graph[n3]
	if ok3 && t.node.directionCount < maxForward {
		neighbors = append(neighbors, Node{n3, t.node.direction, t.node.directionCount + 1})
	}

	return neighbors
}

func heatLoss(graph utils.Grid[int], came_from map[Node]Node, goalTile Tile, start utils.Vertex) int {
	var current = goalTile.node
	var heatLoss = 0
	for current.position != start {
		heatLoss += graph[current.position]
		current = came_from[current]
	}
	return heatLoss
}

func AStar(graph utils.Grid[int], start utils.Vertex, goal utils.Vertex, minForward int, maxForward int) int {
	var frontier = utils.PriorityQueueInit[Tile](lessThan)
	var startTile1 = Tile{Node{start, 1, 0}, 0}
	var startTile2 = Tile{Node{start, 2, 0}, 0}
	frontier.Enqueue(startTile1)
	frontier.Enqueue(startTile2)
	var came_from = make(map[Node]Node)
	var cost_so_far = make(map[Node]int)
	came_from[startTile1.node] = Node{}
	cost_so_far[startTile1.node] = 0
	came_from[startTile2.node] = Node{}
	cost_so_far[startTile2.node] = 0

	for frontier.Count() != 0 {
		var current, _ = frontier.Dequeue()

		if current.node.position == goal && current.node.directionCount > minForward-1 {
			return heatLoss(graph, came_from, current, start)
		}

		for _, next := range current.possibleNeighbors(graph, minForward, maxForward) {
			var new_cost = cost_so_far[current.node] + graph[next.position]
			var _, ok = cost_so_far[next]
			if !ok || new_cost < cost_so_far[next] {
				cost_so_far[next] = new_cost
				var nextTile = Tile{next, new_cost + utils.Manhattan(goal, next.position)}
				frontier.Enqueue(nextTile)
				came_from[next] = current.node
			}
		}
	}
	panic("Goal not found.")
}

func solve(input string, minForward int, maxForward int) int {
	var graph = utils.NewGrid[int]()
	var inputByLines = utils.LineByLine(input)
	for i, line := range inputByLines {
		for j := range line {
			var n, _ = strconv.Atoi(line[j : j+1])
			graph[utils.Vertex{i, j}] = n
		}
	}
	return AStar(graph, utils.Vertex{0, 0}, utils.Vertex{len(inputByLines) - 1, len(inputByLines[0]) - 1}, minForward, maxForward)
}

func partOne(input string) int {
	return solve(input, 0, 3)
}

func partTwo(input string) int {
	return solve(input, 4, 10)
}

func main() {
	var start = time.Now()
	fmt.Println(partOne("17/input.txt"))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(partTwo("17/input.txt"))
	fmt.Println(time.Since(start))
}
