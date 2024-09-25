package main

import "testing"

func TestPartOne(t *testing.T) {
	var actual = exo1("test.txt")
	var expected = 142
	if actual != expected {
		t.Errorf("Part One: actual %d, expected : %d", actual, expected)
	}
}

func TestPartTwo(t *testing.T) {
	var actual = exo2("test2.txt")
	var expected = 281
	if actual != expected {
		t.Errorf("Part Two: actual %d, expected : %d", actual, expected)
	}
}
