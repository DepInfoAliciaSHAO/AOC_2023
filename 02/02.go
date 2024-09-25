package main

import (
	"AOC2023/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func gameToRounds(game string) string {
	return strings.Split(game, ":")[1]
}

func roundsToRound(round string) []string {
	return strings.Split(round, ";")
}

func roundToCubes(set string) []string {
	return strings.Split(set, ",")
}

type Round struct {
	red   int
	blue  int
	green int
}

func (round Round) checkRound(nRed int, nBlue int, nGreen int) bool {
	return round.red <= nRed && round.blue <= nBlue && round.green <= nGreen
}

func firstLetter(cube string) rune {
	for _, c := range cube {
		if unicode.IsLetter(c) {
			return c
		}
	}
	panic("No letter in the string.")
}

func firstNumber(cube string) string {
	for i, c := range cube {
		if unicode.IsLetter(c) {
			return cube[0 : i-1]
		}
	}
	panic("No number in the string.")
}

func checkGame(nRed int, nBlue int, nGreen int, game string) bool {
	var valid = true
	var rounds = gameToRounds(game)
	for _, set := range roundsToRound(rounds) {
		var currentRound = Round{0, 0, 0}
		var cubes = roundToCubes(set)
		for _, cube := range cubes {
			var noSpaceCube = strings.TrimSpace(cube)
			var n, err = strconv.Atoi(firstNumber(noSpaceCube))
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return false
			}
			if firstLetter(noSpaceCube) == 'r' {
				currentRound.red = n
			} else if firstLetter(noSpaceCube) == 'b' {
				currentRound.blue = n
			} else {
				currentRound.green = n
			}
			valid = valid && currentRound.checkRound(nRed, nBlue, nGreen)
		}
	}
	return valid
}

func partOne(input []string, nRed int, nBlue int, nGreen int) int {
	sum := 0
	for i, game := range input {
		if checkGame(nRed, nBlue, nGreen, game) {
			sum += i + 1
		}
	}
	return sum
}

func maxiRound(round1 Round, round2 Round) Round {
	var result = Round{}
	result.red = max(round1.red, round2.red)
	result.green = max(round1.green, round2.green)
	result.blue = max(round1.blue, round2.blue)
	return result
}

func powerGame(game string) int {
	var maxRound = Round{}
	var round = gameToRounds(game)
	for _, set := range roundsToRound(round) {
		var currentRound = Round{0, 0, 0}
		var cubes = roundToCubes(set)
		for _, cube := range cubes {
			var noSpaceCube = strings.TrimSpace(cube)
			var n, err = strconv.Atoi(firstNumber(noSpaceCube))
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return 0
			}
			if firstLetter(noSpaceCube) == 'r' {
				currentRound.red = n
			} else if firstLetter(noSpaceCube) == 'b' {
				currentRound.blue = n
			} else {
				currentRound.green = n
			}
		}
		maxRound = maxiRound(maxRound, currentRound)
	}
	return maxRound.red * maxRound.blue * maxRound.green
}

func partTwo(input []string) int {
	sum := 0
	for _, game := range input {
		sum += powerGame(game)
	}
	return sum
}

func main() {
	var start = time.Now()
	fmt.Println(partOne(utils.LineByLine("02/input.txt"), 12, 14, 13))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(partTwo(utils.LineByLine("02/input.txt")))
	fmt.Println(time.Since(start))
}
