package main

import (
	"AOC2023/utils"
	"fmt"
	"strings"
	"time"
)

func getStartingSequence(line string) []int {
	var stringNumbers = strings.Split(line, " ")
	return utils.StringArrayToIntArray(stringNumbers)
}

func difference(sequence []int) []int {
	var diff = make([]int, len(sequence)-1)
	for i := 0; i < len(sequence)-1; i++ {
		diff[i] = sequence[i+1] - sequence[i]
	}
	return diff
}

func all(slice []int, value int) bool {
	for _, v := range slice {
		if v != value {
			return false
		}
	}
	return true
}

func computeHistory(startingSequence []int) [][]int {
	var history = make([][]int, 0)
	history = append(history, startingSequence)
	var currentSequence = startingSequence
	for !all(currentSequence, 0) {
		currentSequence = difference(currentSequence)
		history = append(history, currentSequence)
	}
	return history
}

func getHistoryValue(history [][]int) int {
	var value = 0
	for i := len(history) - 1; i > -1; i-- {
		value += history[i][len(history[i])-1]
	}
	return value
}

func partOne(input string) int {
	var res = 0
	var inputByLine = utils.LineByLine(input)
	for _, line := range inputByLine {
		var startingSequence = getStartingSequence(line)
		res += getHistoryValue(computeHistory(startingSequence))
	}
	return res
}

func getHistoryValue2(history [][]int) int {
	var value = 0
	for i := len(history) - 1; i > -1; i-- {
		var res = history[i][0] - value
		value = res
	}
	return value
}

func partTwo(input string) int {
	var res = 0
	var inputByLine = utils.LineByLine(input)
	for _, line := range inputByLine {
		var startingSequence = getStartingSequence(line)
		res += getHistoryValue2(computeHistory(startingSequence))
	}
	return res
}

func main() {
	var start = time.Now()
	fmt.Println(partOne("09/input.txt"))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(partTwo("09/input.txt"))
	fmt.Println(time.Since(start))
}
