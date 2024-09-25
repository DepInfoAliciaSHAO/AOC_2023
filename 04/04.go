package main

import (
	"AOC2023/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

/*
Parsing to get rid of the left part of the lines.
*/
func lineToNumbers(line string) string {
	return strings.Split(line, ":")[1]
}

/*
Parsing to separate winning numbers and owned numbers.
*/
func numbersToWinningAndOwned(numbers string) []string {
	return strings.Split(numbers, "|")
}

/*
Solves part 1.
*/
func partOne(input string) int {
	// Open the file
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	var value = 0
	for scanner.Scan() {
		line := scanner.Text()
		//Retrieving winning and owned numbers.
		var WAO = numbersToWinningAndOwned(lineToNumbers(line))
		winningNumbers := utils.StringArrayToIntArray(strings.Fields(WAO[0]))
		ownedNumbers := utils.StringArrayToIntArray(strings.Fields(WAO[1]))
		//For each line, computing the card's value
		var cardValue = 0
		for _, n := range ownedNumbers {
			//Need to initialize the cardValue to one before multiplying it by 2.
			if slices.Contains(winningNumbers, n) {
				if cardValue == 0 {
					cardValue += 1
				} else {
					cardValue *= 2
				}
			}
		}
		value += cardValue
	}
	return value
}

/*
Computes the numbers of won copies from a card.
*/
func numWonCopies(winning []int, owned []int) int {
	var res = 0
	for _, n := range owned {
		if slices.Contains(winning, n) {
			res += 1
		}
	}
	return res
}

/*
Computes the solution for part two.
*/
func countScratchcards(input []string) int {
	var scratchcards = make([]int, len(input))
	//There is at least one copy of each card.
	for i, _ := range scratchcards {
		scratchcards[i] += 1
	}
	for i, line := range input {
		var WAO = numbersToWinningAndOwned(lineToNumbers(line))
		winningNumbers := utils.StringArrayToIntArray(strings.Fields(WAO[0]))
		ownedNumbers := utils.StringArrayToIntArray(strings.Fields(WAO[1]))
		var won = numWonCopies(winningNumbers, ownedNumbers)
		//Adding to the won next cards 1 times the multiplicity of the current card.
		for k := 1; k < won+1; k++ {
			scratchcards[i+k] += scratchcards[i]
		}
	}
	return utils.Sum(scratchcards)
}

func partTwo(input string) int {
	return countScratchcards(utils.LineByLine(input))
}
func main() {
	var start = time.Now()
	fmt.Println(partOne("04/input.txt"))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(partTwo("04/input.txt"))
	fmt.Println(time.Since(start))
}
