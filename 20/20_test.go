package main

import (
	"testing"
)

func TestPartOne1(t *testing.T) {
	var actual = partOne("test1.txt")
	var expected = 32000000
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d.", actual, expected)
	}
}

func TestPartOne2(t *testing.T) {
	var actual = partOne("test2.txt")
	var expected = 11687500
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d.", actual, expected)
	}
}
