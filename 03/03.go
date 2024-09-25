package main

import (
	"AOC2023/utils"
	"fmt"
	"slices"
	"strconv"
	"time"
	"unicode"
)

//////////////
// PART ONE //
//////////////

/*
Returns the result of part one.
*/
func getSum(input []string) int {
	sum := 0
	// Iteration over the lines
	for i, line := range input {
		// Next to symbol is true when a number being built as the line is read is next to a symbol.
		var nextToSymbol = false
		var buildNumber = ""
		// Iteration within a line.
		for j, c := range line {

			//Case where a number has been read and a symbol or a dot is encountered: the number is fully read.
			if buildNumber != "" && !unicode.IsDigit(c) {
				//Conversion to int
				var fullNumber, err = strconv.Atoi(buildNumber)
				if err != nil {
					fmt.Println("Error converting string to int:", err)
					return 0
				}
				//If said number was next to a symbol, it is added to the total sum.
				if nextToSymbol {
					sum += fullNumber
				}

				//Resetting the number builder and the boolean flag for further numbers along the line.
				buildNumber = ""
				nextToSymbol = false

				//Second treated case, the character encountered is a digit.
			} else if unicode.IsDigit(c) {

				//Checking if it's next to a symbol.
				for _, neighbor := range utils.Neighbors8(utils.Vertex{i, j}, len(input), len(input[0])) {
					var c = rune(input[neighbor.X][neighbor.Y])
					if c != '.' && !unicode.IsDigit(c) {
						nextToSymbol = true
					}
				}
				//Building up the number the digit is part of.
				buildNumber += string(c)
			}
		}

		//Necessary lines when a digit is at the end of a line.
		//TODO Compact this, as there is a repetition of code.
		//Note: this condition is sufficient because if the last character is a dot or a symbol, buildNumber is empty.
		if buildNumber != "" {
			var fullNumber, err = strconv.Atoi(buildNumber)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return 0
			}
			if nextToSymbol {
				sum += fullNumber
			}
		}

	}
	return sum
}

func partOne(input string) int {
	return getSum(utils.LineByLine(input))
}

//////////////
// PART TWO //
//////////////

/*
Returns a map where the keys are the coordinates of the stars in the grid,
and their value, a slice of the numbers they are adjacent to.
*/
func getStarMap(input []string) map[utils.Vertex][]int {
	var starMap = make(map[utils.Vertex][]int)
	for i, line := range input {
		//nextToStar is true when a number being built along the line is next to a star symbol.
		var nextToStar = false
		var buildNumber = ""
		//As a number could be next to multiple stars, a slice is used to keep track of them.
		var starNeighbors = make([]utils.Vertex, 0)
		for j, c := range line {
			//Case one: a number has been fully read.
			if buildNumber != "" && !unicode.IsDigit(c) {
				var fullNumber, err = strconv.Atoi(buildNumber)
				if err != nil {
					fmt.Println("Error converting string to int:", err)
					return nil
				}
				//If it is next to a star, the star map is updated accordingly.
				if nextToStar {
					for _, starCoordinates := range starNeighbors {
						starMap[starCoordinates] = append(starMap[starCoordinates], fullNumber)
					}
				}
				//Resetting the memory for next numbers further down the line.
				buildNumber = ""
				nextToStar = false
				starNeighbors = make([]utils.Vertex, 0)

				//Second case: the character is a digit.
			} else if unicode.IsDigit(c) {
				//Checking id it's next to a star.
				for _, neighbor := range utils.Neighbors8(utils.Vertex{i, j}, len(input), len(input[0])) {
					if input[neighbor.X][neighbor.Y] == '*' {
						nextToStar = true
						//Adding it to the list of star neighbors IF it's not already there.
						if !slices.Contains(starNeighbors, neighbor) {
							starNeighbors = append(starNeighbors, neighbor)
						}
					}
				}
				buildNumber += string(c)
			}

			//These lines are necessary if the last character of a line is a digit.
			//TODO again, this is to be compacted.
			//Another condition is added as symbols except * are not checked.
			if buildNumber != "" && j == len(input[0])-1 {
				var fullNumber, err = strconv.Atoi(buildNumber)
				if err != nil {
					fmt.Println("Error converting string to int:", err)
					return nil
				}
				if nextToStar {
					for _, starCoordinates := range starNeighbors {
						starMap[starCoordinates] = append(starMap[starCoordinates], fullNumber)
					}
				}
				buildNumber = ""
				nextToStar = false
				starNeighbors = make([]utils.Vertex, 0)
			}
		}

	}
	return starMap
}

/*
Returns the total gear ratio as defined by part two.
*/
func totalGearRatio(starMap map[utils.Vertex][]int) int {
	gearRatio := 0
	//Gear ratios are added when a star is associated with only two numbers.
	for _, nums := range starMap {
		if len(nums) == 2 {
			gearRatio += nums[0] * nums[1]
		}
	}
	return gearRatio
}

func partTwo(input string) int {
	return totalGearRatio(getStarMap(utils.LineByLine(input)))
}

func main() {
	start := time.Now()
	fmt.Println(partOne("03/input.txt"))
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	start = time.Now()
	fmt.Println(partTwo("03/input.txt"))
	elapsed = time.Since(start)
	fmt.Println(elapsed)
}
