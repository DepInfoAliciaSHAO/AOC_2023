package main

import (
	"AOC2023/utils"
	"fmt"
	"time"
)

////////////
//PART ONE//
////////////

/*
Add the rocks from a given line to the map of the rocks of the platform.
Rocks are placed in the uppermost available free space they can get to.
Free indexes are then updated accordingly, if a rock took the available space,
or if a cube-shaped rock will block the rocks that are south from it.

Do note that the X axis is horizontal, left to right,
and that the Y axis is vertical, up to bottom.
*/
func readLine(line string, rocks map[utils.Vertex]bool, freeIndexes []int, lineIndex int) {
	for i, r := range line {
		switch r {
		case 'O':
			//A rock will first roll to the uppermost free space.
			rocks[utils.Vertex{X: i, Y: freeIndexes[i]}] = true
			//The next available space is thus immediately south from it.
			freeIndexes[i] += 1
		case '#':
			//A cube-shaped rock blocks its own column at its row, the next available space is thus
			//south from it.
			freeIndexes[i] = lineIndex + 1
		case '.':
		default:
			panic("Should not happen.")
		}
	}
}

/*
Computes the total load of a platform. Linear in the number of rocks.
*/
func load(rocks map[utils.Vertex]bool, yDim int) int {
	var totalLoad = 0
	for rock := range rocks {
		totalLoad += yDim - rock.Y
	}
	return totalLoad
}

func partOne(input string) int {
	var inputLines = utils.LineByLine(input)
	var xDim = len(inputLines[0])
	var yDim = len(inputLines)
	var rocks = make(map[utils.Vertex]bool)
	var freeIndexes = make([]int, xDim)
	//Line by line reading to update the state of the platform after the tilt in the right order.
	for i, line := range inputLines {
		readLine(line, rocks, freeIndexes, i)
	}
	return load(rocks, yDim)
}

////////////
//PART TWO//
////////////

type Platform struct {
	cubeShapedRocks map[utils.Vertex]bool
	//As rocks will move, we need an identifier for them, thus the mapping to a unique integer.
	rocks map[utils.Vertex]int
	n     int
}

/*
Checks if two states of the SAME PLATFORM are equal.
For each rock of the first platform state, we check if it's in the second set of rocks.
*/
func equal(p1 Platform, p2 Platform) bool {
	var sameRocks = true
	for rock := range p1.rocks {
		var _, ok = p2.rocks[rock]
		sameRocks = sameRocks && ok
	}
	return sameRocks
}

/*
In-place rotation of the BOARD.
*/
func (p *Platform) clockWiseRotation() {
	var newCSR = make(map[utils.Vertex]bool)
	var newRocks = make(map[utils.Vertex]int)
	for cube := range p.cubeShapedRocks {
		var newPosition = utils.Vertex{X: p.n - 1 - cube.Y, Y: cube.X}
		newCSR[newPosition] = true
	}
	for rock := range p.rocks {
		var newPosition = utils.Vertex{X: p.n - 1 - rock.Y, Y: rock.X}
		newRocks[newPosition] = p.rocks[rock]
	}
	p.rocks = newRocks
	p.cubeShapedRocks = newCSR
}

/*
In place update of the board after a tilt to the North.
Same process as in part one.
*/
func (p *Platform) tiltNorth() {
	var newRocks = make(map[utils.Vertex]int)
	var freeIndexes []int = make([]int, p.n)
	for j := 0; j < p.n; j++ {
		for i := 0; i < p.n; i++ {
			var currentPoint = utils.Vertex{X: i, Y: j}
			var _, okRock = p.rocks[currentPoint]
			var _, okCSR = p.cubeShapedRocks[currentPoint]
			if okRock {
				newRocks[utils.Vertex{X: i, Y: freeIndexes[i]}] = p.rocks[currentPoint]
				freeIndexes[i] += 1
			} else if okCSR {
				freeIndexes[i] = j + 1
			}
		}
	}
	p.rocks = newRocks
}

/*
In place update of the board after a tilt cycle as defined by part two.
*/
func (p *Platform) cycle() {
	p.tiltNorth()
	p.clockWiseRotation()
	p.tiltNorth()
	p.clockWiseRotation()
	p.tiltNorth()
	p.clockWiseRotation()
	p.tiltNorth()
	p.clockWiseRotation()
}

/*
Initialize a platform instance representing the platform from a text input.
The same convention from part one for the axis has been kept.
*/
func initialize(input []string) Platform {
	var CSR = make(map[utils.Vertex]bool)
	var rocks = make(map[utils.Vertex]int)
	var nRocks = 0
	var n = len(input)
	for j, line := range input {
		for i, r := range line {
			switch r {
			case 'O':
				rocks[utils.Vertex{X: i, Y: j}] = nRocks
				nRocks++
			case '#':
				CSR[utils.Vertex{X: i, Y: j}] = true
			}
		}
	}
	return Platform{CSR, rocks, n}
}

/*
Computes the total load of a platform.
Adapted from part one.
*/
func (p *Platform) load() int {
	var totalLoad = 0
	for rock := range p.rocks {
		totalLoad += p.n - rock.Y
	}
	return totalLoad
}

/*
Returns a copy of a platform.
*/
func (p *Platform) copy() Platform {
	var newPlatform = Platform{p.cubeShapedRocks, p.rocks, p.n}
	return newPlatform
}

/*
Computes the index of the first element in the sequence that has the same value as the element at index cycles.
*/
func getIndex(sequence []Platform, firstCheck int, lambda int, cycles int) int {
	//mu will be our offset, assuming the cycle doesn't start at the beginning, mu != 0.
	//mu start at the first checked index.
	var mu = firstCheck
	//while it's in the cyclic sequence, mu is decreased by lambda.
	for mu > 0 && equal(sequence[mu], sequence[firstCheck]) {
		mu -= lambda
	}
	//Let's not forget to re add lambda.
	mu += lambda
	//Let's find the beginning of the cyclic sequence.
	var i = 0
	for mu > 0 && equal(sequence[mu], sequence[firstCheck-i]) {
		mu--
		i++
	}
	mu++
	//If the sequence was fully cyclical, res would be our answer.
	var res = cycles % lambda
	// However, if res is below mu, the smallest index possible is reached by adding lambda
	//until it's above mu.
	for res < mu {
		res += lambda
	}
	return res
}

func partTwo(input string, cycles int) int {
	var platform = initialize(utils.LineByLine(input))
	var sequence = make([]Platform, 0)
	var loadSequence = make([]int, 0)
	loadSequence = append(loadSequence, platform.load())
	sequence = append(sequence, platform.copy())

	var nCycle = 1100     //Total initial cycles to be done.
	var firstCheck = 1000 //Index of first index in the cycle sequence to be checked,
	// has to be big enough for the sequence to stabilize and start its periodic sequence.

	var foundCycle = false
	for !foundCycle {
		var lambda = 0
		for i := 0; i < nCycle; i++ {
			platform.cycle()
			loadSequence = append(loadSequence, platform.load())
			sequence = append(sequence, platform.copy())
			fmt.Printf("\rCycles: %d", len(sequence)-1)
		}
		for i := firstCheck + 1; i < len(sequence); i++ {
			if equal(sequence[firstCheck], sequence[i]) {
				lambda = i - firstCheck
			}
		}
		//Lambda non nul means the sequence has a period.
		if !(lambda == 0) {
			foundCycle = true
			//If the sequence is too small to have a period, additional terms are added.
		} else {
			nCycle += 500
		}
		fmt.Println("")
		return loadSequence[getIndex(sequence, firstCheck, lambda, cycles)]
	}
	panic("Not supposed to happen")
}

func main() {
	var start = time.Now()
	fmt.Println(partOne("14/input.txt"))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(partTwo("14/input.txt", 1000000000))
	fmt.Println(time.Since(start))
}
