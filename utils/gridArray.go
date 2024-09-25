package utils

// Vertex type used to represent coordinates in a 2D grid.
type Vertex struct {
	X int
	Y int
}

type Vertex64 struct {
	X int64
	Y int64
}

/*
From a position X0, Y0 in a grid and the grid's dimension,
returns the positions of valid neighbors of X0, Y0 in the grid.
*/
func neighbors8(vertex Vertex, maxX int, maxY int) []Vertex {
	var neighbors = make([]Vertex, 0)
	//Neighbors are given from left to right, up then down.
	for i := vertex.X - 1; i < vertex.X+2; i++ {
		for j := vertex.Y - 1; j < vertex.Y+2; j++ {
			//The second member of the and condition is there to check if the neighbor is within the grid's bounds.
			if (i != vertex.X || j != vertex.Y) && (i > -1 && i < maxX && j > -1 && j < maxY) {
				neighbors = append(neighbors, Vertex{i, j})
			}
		}
	}
	return neighbors
}
