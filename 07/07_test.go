package main

import "testing"

func TestPartOne(t *testing.T) {
	var actual = partOne("test.txt")
	var expected = 6440
	if actual != expected {
		t.Errorf("Score = %d; want %d", actual, expected)
	}
}

func TestPartOne2(t *testing.T) {
	var actual = partOne("test2.txt")
	var expected = 1343
	if actual != expected {
		t.Errorf("Score = %d; want %d", actual, expected)
	}
}

func TestPartTwo(t *testing.T) {
	var actual = partTwo("test.txt")
	var expected = 5905
	if actual != expected {
		t.Errorf("Score = %d; want %d", actual, expected)
	}
}

func TestPartTwo2(t *testing.T) {
	var actual = partTwo("test2.txt")
	var expected = 1369
	if actual != expected {
		t.Errorf("Score = %d; want %d", actual, expected)
	}
}
