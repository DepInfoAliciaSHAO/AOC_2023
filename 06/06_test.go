package main

import "testing"

func TestPartOne(t *testing.T) {
	var actual = partOne("test.txt")
	var expected = 288
	if actual != expected {
		t.Errorf("Product = %d; want %d", actual, expected)
	}
}

func TestPartTwo(t *testing.T) {
	var actual = partTwo("test.txt")
	var expected = 71503
	if actual != expected {
		t.Errorf("Ways to beat record = %d; want %d", actual, expected)
	}
}
