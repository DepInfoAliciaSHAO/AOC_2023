package utils

import (
	"fmt"
	"strconv"
)

/*
Sums the elements of an integer slice.
*/
func Sum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

/*
Converts a string array of string representations of int to the associated int array.
*/
func StringArrayToIntArray(numbers []string) []int {
	var res []int = make([]int, 0)
	for _, str := range numbers {
		var n, err = strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error converting to int:", err)
			return res
		}
		res = append(res, n)
	}
	return res
}

func arrayToInt(array []string) int {
	var str string = ""
	for _, s := range array {
		str += s
	}
	var n, _ = strconv.Atoi(str)
	return n
}

func MinimumArray(array []int) int {
	var mini = array[0]
	for _, v := range array {
		if mini > v {
			mini = v
		}
	}
	return mini
}

func MaximumArray(array []int) int {
	var max = array[0]
	for _, v := range array {
		if max < v {
			max = v
		}
	}
	return max
}

func Every(array []int, value int) {
	for i := range array {
		array[i] = value
	}
}
