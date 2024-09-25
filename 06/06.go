package main

import (
	"AOC2023/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func getData(input string) ([]int, []int) {

	// Open the file
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	var time = make([]int, 0)
	var distance = make([]int, 0)
	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Time: ") {
			var numbers, _ = strings.CutPrefix(line, "Time: ")
			var numberArray = strings.Split(strings.TrimSpace(numbers), " ")
			time = utils.StringArrayToIntArray(numberArray)
		} else {
			var numbers, _ = strings.CutPrefix(line, "Distance: ")
			var numberArray = strings.Split(strings.TrimSpace(numbers), " ")
			distance = utils.StringArrayToIntArray(numberArray)
		}
	}
	return time, distance
}

func recordBeats(time int, distance int) int {
	var nRecordBeat = 0
	for i := 1; i < time; i++ {
		var speed = i
		var timeTravelled = time - i
		var distanceTravelled = speed * timeTravelled
		if distanceTravelled > distance {
			nRecordBeat += 1
		}
	}
	return nRecordBeat
}

func getProduct(time []int, distance []int) int {
	var product = 1
	for i := 0; i < len(time); i++ {
		product *= recordBeats(time[i], distance[i])
	}
	return product
}

func partOne(input string) int {
	return getProduct(getData(input))
}

//DT = i * (time - i) > distance
//f(i) = - i^2 + i * time - distance > 0
// delta = time^2 - 4 * distance
// x*_ = time + sqrt(delta) /2
// x*+ = time - sqrt(delta) /2

func solution(time int, distance int) int {
	var delta = time*time - 4*distance
	var xPlus = (time + int(math.Sqrt(float64(delta)))) / 2
	var xMinus = (time - int(math.Sqrt(float64(delta)))) / 2
	return xPlus - xMinus
}

func arrayToInt(array []string) int {
	var str string = ""
	for _, s := range array {
		str += s
	}
	var n, _ = strconv.Atoi(str)
	return n
}

func getNumbers(input string) (int, int) {

	// Open the file
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, 0
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	var time = 0
	var distance = 0
	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Time: ") {
			var numbers, _ = strings.CutPrefix(line, "Time: ")
			var numberArray = strings.Split(strings.TrimSpace(numbers), " ")
			time = arrayToInt(numberArray)
		} else {
			var numbers, _ = strings.CutPrefix(line, "Distance: ")
			var numberArray = strings.Split(strings.TrimSpace(numbers), " ")
			distance = arrayToInt(numberArray)
		}
	}
	return time, distance
}

func partTwo(input string) int {
	return solution(getNumbers(input))
}

func main() {
	var start = time.Now()
	fmt.Println(partOne("06/input.txt"))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(partTwo("06/input.txt"))
	fmt.Println(time.Since(start))
}
