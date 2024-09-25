package main

import (
	"AOC2023/utils"
	"fmt"
	"time"
)

/*
Returns the coordinates of the galaxy in the unexpanded map and
lineGalaxies such that if there is at least a galaxy on line i, lineGalaxies[i] = 0
and if there is no galaxy on line i, lineGalaxies[i] = 1.
Same thing for columnGalaxies with columns.
*/
func getMapImage(input []string) ([]utils.Vertex, []int, []int) {
	var galaxies = make([]utils.Vertex, 0)
	//Default: no galaxy on each line and column
	var lineGalaxies = make([]int, len(input))
	utils.Every(lineGalaxies, 1)
	var columnGalaxies = make([]int, len(input[0]))
	utils.Every(columnGalaxies, 1)
	for i, line := range input {
		for j, r := range line {
			if rune(r) == '#' {
				//There is a galaxy on line i and in column j
				galaxies = append(galaxies, utils.Vertex{X: i, Y: j})
				lineGalaxies[i] = 0
				columnGalaxies[j] = 0
			}
		}
	}
	return galaxies, lineGalaxies, columnGalaxies
}

type GalaxyMap struct {
	galaxies       []utils.Vertex
	lineGalaxies   []int
	columnGalaxies []int
}

/*
From a galaxy map, computes the actual distance from a galaxy A to a galaxy B.
Expansion is the expansion factor.
The actual distance is the differance along the X and Y axis of A and B's coordinates.
In the unexpanded map, when each line or column with no galaxies is crossed, expansion is added to distance
to take it into account.
*/
func (gm GalaxyMap) distance(galaxyA int, galaxyB int, expansion int) int {
	var distance = 0
	var minX, maxX = 0, 0
	var minY, maxY = 0, 0
	// Trivial case where A = B
	if galaxyA == galaxyB {
		return 0
	} else {
		// maxZ and minZ are needed for X and Y to check if the lines/columns in between them have to be expanded.
		maxX = utils.Max(gm.galaxies[galaxyA].X, gm.galaxies[galaxyB].X)
		minX = utils.Min(gm.galaxies[galaxyA].X, gm.galaxies[galaxyB].X)
		maxY = utils.Max(gm.galaxies[galaxyA].Y, gm.galaxies[galaxyB].Y)
		minY = utils.Min(gm.galaxies[galaxyA].Y, gm.galaxies[galaxyB].Y)

		// Actual distance
		distance += maxX - minX
		distance += maxY - minY

		// Expansion (one line or column has already been taken into account, thus the -1)
		for i := minX; i < maxX; i++ {
			distance += (expansion - 1) * gm.lineGalaxies[i]
		}

		for j := minY; j < maxY; j++ {
			distance += (expansion - 1) * gm.columnGalaxies[j]
		}
	}

	return distance
}

/*
Computes the total minimal distances between the galaxies of a given galaxy map according to a given expansion factor.
*/
func (gm GalaxyMap) computeDistance(expansion int) int {
	var totalDistance = 0
	//Triangular sum to avoid computing the distances twice.
	for i := 0; i < len(gm.galaxies); i++ {
		for j := 0; j < i; j++ {
			totalDistance += gm.distance(i, j, expansion)
		}
	}
	return totalDistance
}
func partOne(input string) int {
	var galaxies, line, colum = getMapImage(utils.LineByLine(input))
	var gm = GalaxyMap{galaxies, line, colum}
	return gm.computeDistance(2)
}

func partTwo(input string, expansion int) int {
	var galaxies, line, colum = getMapImage(utils.LineByLine(input))
	var gm = GalaxyMap{galaxies, line, colum}
	return gm.computeDistance(expansion)
}

func main() {
	var start = time.Now()
	fmt.Println(partOne("11/input.txt"))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(partTwo("11/input.txt", 1000000))
	fmt.Println(time.Since(start))
}
