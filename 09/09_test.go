package main

import "testing"

func TestPartOne(t *testing.T) {
	var actual = partOne("test.txt")
	var expected = 114
	if actual != expected {
		t.Errorf("Actual: %d, expected: %d.", actual, expected)
	}
}

func TestPartTwo(t *testing.T) {
	var actual = partTwo("test.txt")
	var expected = 2
	if actual != expected {
		t.Errorf("Actual: %d, expected: %d.", actual, expected)
	}
}
