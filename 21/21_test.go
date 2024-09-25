package main

import "testing"

func TestPartOne1(t *testing.T) {
	var actual = partOne("test.txt", 6)
	var expected = 16
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d.", actual, expected)
	}
}
