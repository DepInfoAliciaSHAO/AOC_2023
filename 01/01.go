package main

import (
	"AOC2023/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

/*
Returns the relevant numbers for each line in part 1 in a slice.
*/
func getNumbers1(slice []string) []int {
	var numbers []int = make([]int, len(slice))
	// Iteration over the lines
	for i, line := range slice {
		var first = -1
		var last = -1
		// Iteration over a line
		for j, c := range line {
			// If c is a number
			if unicode.IsNumber(c) {
				var n, _ = strconv.Atoi(line[j : j+1])
				first, last = getDigits(first, last, n)
			}
		}

		//Case where there's only one number
		if last == -1 {
			last = first
		}

		// Conversion
		var result = 10*first + last
		numbers[i] = result
	}
	return numbers
}

/*
Part 1: Reading the input.txt file, retrieving relevant numbers from each line and summing them.
*/
func exo1(input string) int {
	var res = utils.Sum(getNumbers1(utils.LineByLine(input)))
	return res
}

/*
Assigns the first or last place to a string depending on if there's already a number in first.
write is a digit string, first and last are the current digit strings.
*/
func getDigits(first int, last int, n int) (int, int) {
	if first == -1 {
		//First digit
		first = n
	} else {
		//Last digit.
		last = n
	}
	return first, last
}

var SPELLED_OUT_DIGIT = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

/*
Checks written number prefixes in a word.
*/
func checkPrefix(first int, last int, check string) (int, int) {
	for i, word := range SPELLED_OUT_DIGIT {
		if strings.HasPrefix(check, word) {
			return getDigits(first, last, i)
		}
	}
	return first, last
}

/*
Getting for each line its associated number as defined by part 2.
*/
func getNumbers2(slice []string) []int {
	var numbers = make([]int, len(slice))
	// Iteration over the lines
	for i, line := range slice {
		var first = -1
		var last = -1
		// Iteration over a line
		for j, c := range line {
			// If c is a digit
			if unicode.IsNumber(c) {
				var n, _ = strconv.Atoi(line[j : j+1])
				first, last = getDigits(first, last, n)
			} else {
				//If c is a letter:
				//Let us brute force if it's the starting letter of a written digit.
				var check = line[j:]
				first, last = checkPrefix(first, last, check)
			}
		}

		//Case where there's only one number on the line.
		if last == -1 {
			last = first
		}

		// Conversion.
		var result = 10*first + last
		numbers[i] = result
	}
	return numbers
}

func exo2(input string) int {
	var res = utils.Sum(getNumbers2(utils.LineByLine(input)))
	fmt.Println(res)
	return res
}

func main() {
	start := time.Now()
	fmt.Println(exo1("01/input.txt"))
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	start = time.Now()
	fmt.Println(exo2("01/input.txt"))
	elapsed = time.Since(start)
	fmt.Println(elapsed)
}
