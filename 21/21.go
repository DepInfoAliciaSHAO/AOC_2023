package main

import (
	"AOC2023/utils"
	"fmt"
	"time"
)

func getStartAndMap(input []string) (utils.Vertex, utils.Grid[string]) {
	var start = utils.Vertex{}
	var fullMap = utils.NewGrid[string]()
	for i, line := range input {
		for j, r := range line {
			var currentVertex = utils.Vertex{i, j}
			switch r {
			case '#':
				fullMap[currentVertex] = "#"
			case 'S':

				start = currentVertex
				fullMap[currentVertex] = "."
			case '.':
				fullMap[currentVertex] = "."
			default:
				panic("Unknown rune.")
			}
		}
	}
	return start, fullMap
}

func cardinal(set map[utils.Vertex]bool) int {
	var cardinal = 0
	for element := range set {
		if set[element] {
			cardinal += 1
		}
	}
	return cardinal
}

func copySet(set map[utils.Vertex]bool) map[utils.Vertex]bool {
	var newSet = make(map[utils.Vertex]bool)
	for element := range set {
		if set[element] {
			newSet[element] = true
		}
	}
	return newSet
}

func gardenPots(maxSteps int, fullMap utils.Grid[string], currentPositions map[utils.Vertex]bool, iteration int) (int, map[utils.Vertex]bool) {
	if iteration == maxSteps {
		return cardinal(currentPositions), currentPositions
	} else {
		var newPositions = copySet(currentPositions)
		for position := range currentPositions {
			if currentPositions[position] {
				newPositions[position] = false
				var neighbors = position.Neighbors4()
				for _, neighbor := range neighbors {
					var _, inMap = fullMap[neighbor]
					if fullMap[neighbor] != "#" && inMap {
						newPositions[neighbor] = true
					}
				}
			}
		}
		return gardenPots(maxSteps, fullMap, newPositions, iteration+1)
	}
}

func partOne(input string, maxSteps int) int {
	var start, fullMap = getStartAndMap(utils.LineByLine(input))
	var currentPositons = make(map[utils.Vertex]bool)
	currentPositons[start] = true
	var res, _ = gardenPots(maxSteps, fullMap, currentPositons, 0)
	return res
}

func parity(fullMap utils.Grid[string], dim int) (map[utils.Vertex]bool, map[utils.Vertex]bool, int, int) {
	var currentPositons = make(map[utils.Vertex]bool)
	currentPositons[utils.Vertex{X: dim / 2, Y: dim / 2}] = true
	var nEven, even = gardenPots(dim+1, fullMap, currentPositons, 0)
	var nOdd, odd = gardenPots(dim, fullMap, currentPositons, 0)
	return even, odd, nEven, nOdd
}

// 26501365 = 202300 * 131 + 65
// 131 = 2 * 65 + 1
// 625102149288785 too low
// 626166222608250 too high
func main() {
	var start = time.Now()
	fmt.Println(partOne("21/input.txt", 64))
	fmt.Println(time.Since(start))
}
