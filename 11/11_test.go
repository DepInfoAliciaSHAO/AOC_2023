package main

import "testing"

func TestPartOne(t *testing.T) {
	var actual = partOne("test.txt")
	var expected = 374
	if actual != expected {
		t.Errorf("Actual: %d, expected: %d.", actual, expected)
	}
}

func TestPartTwo100(t *testing.T) {
	var actual = partTwo("test.txt", 100)
	var expected = 8410
	if actual != expected {
		t.Errorf("Actual: %d, expected: %d.", actual, expected)
	}
}
