package main

import "testing"

func TestPartOne(t *testing.T) {
	var actual = partOne("test.txt")
	var expected = 102
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d.", actual, expected)
	}
}

func TestPartOne2(t *testing.T) {
	var actual = partOne("test2.txt")
	var expected = 8
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d.", actual, expected)
	}
}

func TestPartOne3(t *testing.T) {
	var actual = partOne("test3.txt")
	if actual == 8 {
		t.Errorf("Actual: %d", actual)
	}
}

func TestPartTwo(t *testing.T) {
	var actual = partTwo("test.txt")
	var expected = 94
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d.", actual, expected)
	}
}

func TestPartTwo2(t *testing.T) {
	var actual = partTwo("test4.txt")
	var expected = 71
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d.", actual, expected)
	}
}
