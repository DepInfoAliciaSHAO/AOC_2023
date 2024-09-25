package main

import "testing"

func TestPartOne(t *testing.T) {
	var actual = partOne("test.txt")
	var expected = 8
	if actual != expected {
		t.Errorf("Actual: %d, expected: %d.", actual, expected)
	}
}

func TestPartTwo4(t *testing.T) {
	var actual = partTwo("test3.txt")
	var expected = 4
	if actual != expected {
		t.Errorf("Actual: %d, expected: %d.", actual, expected)
	}
}

func TestPartTwo10(t *testing.T) {
	var actual = partTwo("test2.txt")
	var expected = 10
	if actual != expected {
		t.Errorf("Actual: %d, expected: %d.", actual, expected)
	}
}

func TestPartTwo8(t *testing.T) {
	var actual = partTwo("test4.txt")
	var expected = 8
	if actual != expected {
		t.Errorf("Actual: %d, expected: %d.", actual, expected)
	}
}
