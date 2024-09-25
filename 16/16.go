package main

import (
	"AOC2023/utils"
	"fmt"
	"time"
)

//North: 0, West: 1, East: 2, South: 3

/*
From a vertex and a direction, returns the next vertex of the traversal.
*/
func nextVertex(v utils.Vertex, direction int) utils.Vertex {
	switch direction {
	case 0:
		return v.UpNeighbor()
	case 1:
		return v.LeftNeighbor()
	case 2:
		return v.RightNeighbor()
	case 3:
		return v.DownNeighbor()
	default:
		panic("Not supposed to happen.")
	}
}

/*
Returns the new direction of the traversal after encountering a / mirror.
*/
func slashTurn(direction int) int {
	switch direction {
	case 0:
		return 2
	case 2:
		return 0
	case 1:
		return 3
	case 3:
		return 1
	default:
		panic("Not supposed to happen.")
	}
}

/*
Returns the new direction of the traversal after encountering a \ mirror.
*/
func backslashTurn(direction int) int {
	switch direction {
	case 0:
		return 1
	case 1:
		return 0
	case 2:
		return 3
	case 3:
		return 2
	default:
		panic("Not supposed to happen.")
	}
}

/*
Returns the two new directions of the traversal after encountering a - mirror.
If the mirror was encountered did not cause a split, both returned directions are the original direction.
The third boolean value indicates whether the beam was split in two.
*/
func dashTurn(direction int) (int, int, bool) {
	switch direction {
	case 0:
		return 1, 2, true
	case 1:
		return 1, 1, false
	case 2:
		return 2, 2, false
	case 3:
		return 1, 2, true
	default:
		panic("Not supposed to happen.")
	}
}

/*
Returns the two new directions of the traversal after encountering a | mirror.
If the mirror was encountered did not cause a split, both returned directions are the original direction.
The third boolean value indicates whether the beam was split in two.
*/
func stickTurn(direction int) (int, int, bool) {
	switch direction {
	case 0:
		return 0, 0, false
	case 1:
		return 0, 3, true
	case 2:
		return 0, 3, true
	case 3:
		return 3, 3, false
	default:
		panic("Not supposed to happen.")
	}
}

// Exploration /* A type that encapsulates the ways to explore a tile (position + direction).
type Exploration struct {
	direction int
	point     utils.Vertex
}

/*
Does a traversal of the mirror map from a starting point recursively.
mirrorMap is the input grid.

energizedMap represents the lit tiles and is updated throughout the traversal.
explored is the set of explorations.
starting is the starting vertex of the traversal.
lit is the number of tiles that are already lit up. At first, there is none.
direction is the starting direction of the traversal.

Returns the number of lit tiles by the traversal.
*/
func traversal(mirrorMap utils.Grid[string], energizedMap utils.Grid[int], explored map[Exploration]bool, starting utils.Vertex, lit int, direction int) int {
	var symbol, inMap = mirrorMap[starting]
	//If the starting point is not within the mirror map, it means the beam reached the end of its journey.
	if !inMap {
		return lit
	} else {
		var currentExploration = Exploration{direction, starting}
		var hasBeenExplored = explored[currentExploration]
		if hasBeenExplored {
			//If the current vertex has been explored in the same direction, the beam loops upon itself.
			//and cannot light up any more tiles: the recursion is stopped.
			return lit
		} else {
			explored[currentExploration] = true
		}
		//Perhaps the tile the light beam is at is already lit up.
		//In this case, the number of lit up tiles doesn't change.
		var _, energized = energizedMap[starting]
		if !energized {
			energizedMap[starting] = 1
			lit += 1
		}

		//Continuing exploration
		switch symbol {

		//Same direction
		case ".":
			return traversal(mirrorMap, energizedMap, explored, nextVertex(starting, direction), lit, direction)

			// Direction changes
		case "/":
			var newDirection = slashTurn(direction)
			return traversal(mirrorMap, energizedMap, explored, nextVertex(starting, newDirection), lit, newDirection)

		case "\\":
			var newDirection = backslashTurn(direction)
			return traversal(mirrorMap, energizedMap, explored, nextVertex(starting, newDirection), lit, newDirection)

			//If the beam is eventually split, two light beams are considered:
			//The original one continues to count up the lit tiles in one direction.
			//A new one, i.e. that hasn't lit up any tiles yet start from the other direction.
		case "-":
			var new1, new2, split = dashTurn(direction)
			if split {
				return traversal(mirrorMap, energizedMap, explored, nextVertex(starting, new1), lit, new1) +
					traversal(mirrorMap, energizedMap, explored, nextVertex(starting, new2), 0, new2)
			} else {
				return traversal(mirrorMap, energizedMap, explored, nextVertex(starting, new1), lit, new1)
			}

		case "|":
			var new1, new2, split = stickTurn(direction)
			if split {
				return traversal(mirrorMap, energizedMap, explored, nextVertex(starting, new1), lit, new1) +
					traversal(mirrorMap, energizedMap, explored, nextVertex(starting, new2), 0, new2)
			} else {
				return traversal(mirrorMap, energizedMap, explored, nextVertex(starting, new1), lit, new1)
			}
		}
	}
	panic("Not supposed to happen.")
}

/*
Creates the mirror map from the input given line by line in a string array.
*/
func makeMirrorMap(input []string) utils.Grid[string] {
	var mirrorMap = utils.NewGrid[string]()
	for i, line := range input {
		for j := range line {
			mirrorMap[utils.Vertex{X: i, Y: j}] = line[j : j+1]
		}
	}
	return mirrorMap
}

/*
Solves part one.
*/
func partOne(input string) int {
	var mirrorMap = makeMirrorMap(utils.LineByLine(input))
	return doTraversal(mirrorMap, utils.Vertex{}, 2)
}

/*
Does a traversal of a mirrorMap from a given starting point, in a given starting direction.
Returns the number of lit up tiles along the traversal.
*/
func doTraversal(mirrorMap utils.Grid[string], startingPoint utils.Vertex, startingDirection int) int {
	var energizedMap = utils.NewGrid[int]()
	var explored = make(map[Exploration]bool)
	return traversal(mirrorMap, energizedMap, explored, startingPoint, 0, startingDirection)
}

/*
Solves part two naïvely: each possibility is checked.
*/
func partTwo(input string) int {
	var inputLineByLine = utils.LineByLine(input)
	var maxX = len(inputLineByLine)
	var maxY = len(inputLineByLine[0])
	var mirrorMap = makeMirrorMap(utils.LineByLine(input))
	var maxEnergy = 0

	// Northern traversals
	fmt.Println("Starting northern edge.")
	for j := 0; j < maxY; j++ {
		fmt.Printf("\rTraversals: %d/%d", j+1, maxY)
		var startingPoint = utils.Vertex{Y: j}
		var startingDirection = 3
		var energy = doTraversal(mirrorMap, startingPoint, startingDirection)
		maxEnergy = max(maxEnergy, energy)
	}
	// Western traversals
	fmt.Println("\nStarting western edge.")
	for i := 0; i < maxX; i++ {
		fmt.Printf("\rTraversals: %d/%d", i+1, maxX)
		var startingPoint = utils.Vertex{X: i}
		var startingDirection = 2
		var energy = doTraversal(mirrorMap, startingPoint, startingDirection)
		maxEnergy = max(maxEnergy, energy)
	}
	// Eastern traversals
	fmt.Println("\nStarting eastern edge.")
	for i := 0; i < maxX; i++ {
		fmt.Printf("\rTraversals: %d/%d", i+1, maxX)
		var startingPoint = utils.Vertex{X: i, Y: maxY - 1}
		var startingDirection = 1
		var energy = doTraversal(mirrorMap, startingPoint, startingDirection)
		maxEnergy = max(maxEnergy, energy)
	}

	// Southern traversals
	fmt.Println("\nStarting southern edge.")
	for j := 0; j < maxY; j++ {
		fmt.Printf("\rTraversals: %d/%d", j+1, maxY)
		var startingPoint = utils.Vertex{X: maxX - 1, Y: j}
		var startingDirection = 0
		var energy = doTraversal(mirrorMap, startingPoint, startingDirection)
		maxEnergy = max(maxEnergy, energy)
	}
	fmt.Println("")
	return maxEnergy
}
func main() {
	var start = time.Now()
	fmt.Println(partOne("16/input.txt"))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(partTwo("16/input.txt"))
	fmt.Println(time.Since(start))
}
