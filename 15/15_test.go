package main

import "testing"

func TestPartOne(t *testing.T) {
	var actual = partOne("test.txt")
	var expected = 1320
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d.", actual, expected)
	}
}

func TestPartTwo(t *testing.T) {
	var actual = partTwo("test.txt")
	var expected = 145
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d.", actual, expected)
	}
}
