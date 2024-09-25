package main

import "testing"

func TestPartOne1(t *testing.T) {
	var expected = 4361
	var actual = partOne("test.txt")
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d", actual, expected)
	}
}

func TestPartOne2(t *testing.T) {
	var expected = 925
	var actual = partOne("test2.txt")
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d", actual, expected)
	}
}

func TestPartTwo1(t *testing.T) {
	var expected = 4361
	var actual = partOne("test.txt")
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d", actual, expected)
	}
}

func TestPartTwo2(t *testing.T) {
	var expected = 6756
	var actual = partTwo("test2.txt")
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d", actual, expected)
	}
}
