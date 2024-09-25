package main

import (
	"AOC2023/utils"
	"testing"
)

func TestPartOne(t *testing.T) {
	var actual = partOne("test.txt")
	var expected = 136
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d.", actual, expected)
	}
}

func TestFloyd(t *testing.T) {
	var sequence = []int{1, 5, 6, 4, 3, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3}
	var actualLambda, actualMu = utils.FloydCycleDetection(sequence)
	var expectedLambda, expectedMu = 3, 5
	if actualMu != expectedMu || actualLambda != expectedLambda {
		t.Errorf("Actual mu, lambda: %d, %d, expected mu, lambda : %d, %d.", actualMu, actualLambda, expectedMu, expectedLambda)
	}
}

func TestPartTwo(t *testing.T) {
	var actual = partTwo("test.txt", 1000000000)
	var expected = 64
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d.", actual, expected)
	}
}
