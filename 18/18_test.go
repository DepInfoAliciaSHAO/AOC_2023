package main

import (
	"AOC2023/utils"
	"testing"
)

func TestShoelace(t *testing.T) {
	var vertices = []utils.Vertex{utils.Vertex{}, utils.Vertex{3, 0}, utils.Vertex{0, 4}}
	var actual = utils.Shoelace(vertices)
	var expected = 6
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d.", actual, expected)
	}
}

func TestPartOne(t *testing.T) {
	var actual = partOne("test.txt")
	var expected = 62
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d.", actual, expected)
	}
}

func TestPartTwo(t *testing.T) {
	var actual = partTwo("test.txt")
	var expected = 952408144115
	if actual != expected {
		t.Errorf("Actual: %d, expected : %d.", actual, expected)
	}
}
