package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	var actual = partOne("test.txt")
	var expected = 13
	if actual != expected {
		t.Errorf("PartOne = %d; want %d", actual, expected)
	}
}

func TestPartTwo(t *testing.T) {
	var actual = partTwo("test.txt")
	var expected = 30
	if actual != expected {
		t.Errorf("PartTwo = %d; want %d", actual, expected)
	}
}
