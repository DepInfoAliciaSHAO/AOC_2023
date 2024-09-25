package main

import (
	"AOC2023/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func getSteps(input string) []string {
	return strings.Split(utils.LineByLine(input)[0], ",")
}

func hashAlgorithm(step string) int {
	var res = 0
	for _, r := range step {
		res += int(r)
		res *= 17
		res = res % 256
	}
	return res
}

func computeSteps(steps []string) []int {
	var stepRes = make([]int, 0)
	for _, step := range steps {
		stepRes = append(stepRes, hashAlgorithm(step))
	}
	return stepRes
}

func partOne(input string) int {
	var steps = getSteps(input)
	var stepHash = computeSteps(steps)
	return utils.Sum(stepHash)
}

type Lens struct {
	id          string
	focalLength int
}

type Boxes = map[int][]Lens

func makeBoxes() map[int][]Lens {
	return make(map[int][]Lens)
}

func add(b Boxes, hash int, lens Lens) {
	var found = false
	var lensList = b[hash]
	for i, l := range lensList {
		if l.id == lens.id {
			found = true
			lensList[i] = lens
			break
		}
	}
	if !found {
		b[hash] = append(b[hash], lens)
	}
}
func computeEqual(b Boxes, step string) {
	var components = strings.Split(step, "=")
	var id = components[0]
	var focalLength, _ = strconv.Atoi(components[1])
	var currentLens = Lens{id, focalLength}
	var hash = hashAlgorithm(currentLens.id)
	add(b, hash, currentLens)
}

func remove(b Boxes, hash int, id string) {
	var lensList = b[hash]
	var stopIndex = 0
	var found = false
	for i, l := range lensList {
		if l.id == id {
			stopIndex = i
			found = true
			break
		}
	}
	if found {
		var newList = append(lensList[:stopIndex], lensList[stopIndex+1:]...)
		b[hash] = newList
	}
}

func computeMinus(b Boxes, step string) {
	var components = strings.Split(step, "-")
	var id = components[0]
	var hash = hashAlgorithm(id)
	remove(b, hash, id)
}

func computeStep(b Boxes, step string) {
	if strings.Contains(step, "=") {
		computeEqual(b, step)
	} else {
		computeMinus(b, step)
	}
}

func focusingPower(b Boxes) int {
	var focusingPowerPerLens = make([]int, 0)
	for box := range b {
		for i, lens := range b[box] {
			focusingPowerPerLens = append(focusingPowerPerLens, (box+1)*(i+1)*lens.focalLength)
		}
	}
	return utils.Sum(focusingPowerPerLens)
}

func partTwo(input string) int {
	var steps = getSteps(input)
	var boxes = makeBoxes()
	for _, step := range steps {
		computeStep(boxes, step)
	}
	return focusingPower(boxes)
}

func main() {
	var start = time.Now()
	fmt.Println(partOne("15/input.txt"))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(partTwo("15/input.txt"))
	fmt.Println(time.Since(start))
}
