package main

import "testing"

func TestPartOne(t *testing.T) {
	var actual = partOne(lineByLine("test.txt"), 12, 14, 13)
	var expected = 8
	if actual != expected {
		t.Errorf("Actual: %d, Expected :%d", actual, expected)
	}
}

func TestPartTwo(t *testing.T) {
	var actual = partTwo(lineByLine("test.txt"))
	var expected = 2286
	if actual != expected {
		t.Errorf("Actual: %d, Expected :%d", actual, expected)
	}
}
