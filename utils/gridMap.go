package utils

type Grid[T comparable] map[Vertex]T

func NewGrid[T comparable]() Grid[T] {
	return make(Grid[T])
}

func Manhattan(v1 Vertex, v2 Vertex) int {
	return Abs(v2.X-v1.X) + Abs(v2.Y-v1.Y)
}

func Shoelace(vertices []Vertex) int {
	var sum = 0
	for i := range vertices {
		var j = i + 1
		if i == len(vertices)-1 {
			j = 0
		}
		sum += (vertices[i].Y + vertices[j].Y) * (vertices[i].X - vertices[j].X)
	}
	return Abs(sum) / 2
}

func Shoelace64(vertices []Vertex64) int64 {
	var crossProduct1 = int64(0)
	var crossProduct2 = int64(0)
	for i := range vertices {
		var j = i + 1
		if i == len(vertices)-1 {
			j = 0
		}
		crossProduct1 += vertices[i].X * vertices[j].Y
		crossProduct2 += vertices[i].Y * vertices[j].X
	}
	return Abs64(crossProduct1-crossProduct2) / 2
}

func Pick(area int, length int) int {
	return area + length/2 + 1
}

func Pick64(area int64, length int64) int64 {
	return area + length/2 + 1
}

func (v Vertex) neighbors8() []Vertex {
	var neighbors = make([]Vertex, 0)
	//Neighbors are given from left to right, up then down.
	for i := v.X - 1; i < v.X+2; i++ {
		for j := v.Y - 1; j < v.Y+2; j++ {
			//The second member of the and condition is there to check if the neighbor is within the grid's bounds.
			if i != v.X || j != v.Y {
				neighbors = append(neighbors, Vertex{i, j})
			}
		}
	}
	return neighbors
}

func (v Vertex) Neighbors4() []Vertex {
	return []Vertex{
		{X: v.X, Y: v.Y - 1},
		{X: v.X, Y: v.Y + 1},
		{X: v.X - 1, Y: v.Y},
		{X: v.X + 1, Y: v.Y},
	}
}

func (v Vertex) UpNeighbor() Vertex {
	return Vertex{X: v.X - 1, Y: v.Y}
}

func (v Vertex) LeftNeighbor() Vertex {
	return Vertex{X: v.X, Y: v.Y - 1}
}

func (v Vertex) RightNeighbor() Vertex {
	return Vertex{X: v.X, Y: v.Y + 1}
}

func (v Vertex) DownNeighbor() Vertex {
	return Vertex{X: v.X + 1, Y: v.Y}
}

func (v Vertex) Move(deltaX int, deltaY int) Vertex {
	return Vertex{X: v.X + deltaX, Y: v.Y + deltaY}
}

func (v Vertex64) Move64(deltaX int64, deltaY int64) Vertex64 {
	return Vertex64{X: v.X + deltaX, Y: v.Y + deltaY}
}

func (g Grid[T]) add(element T, position Vertex) {
	g[position] = element
}

func (g Grid[T]) read(position Vertex) T {
	return g[position]
}

func (g Grid[T]) has(element T) bool {
	for _, value := range g {
		if value == element {
			return true
		}
	}
	return false
}

func (g Grid[T]) addRow(element T, start Vertex, nAdded int) {
	for j := start.Y; j < nAdded; j++ {
		g.add(element, Vertex{start.X, j})
	}
}

func (g Grid[T]) addColumn(element T, start Vertex, nAdded int) {
	for i := start.X; i < nAdded; i++ {
		g.add(element, Vertex{i, start.Y})
	}
}

func (g Grid[T]) addArrayHorizontal(array []T, start Vertex) {
	for j, element := range array {
		g.add(element, Vertex{start.X, start.Y + j})
	}
}

func (g Grid[T]) addArrayVertical(array []T, start Vertex) {
	for i, element := range array {
		g.add(element, Vertex{start.X + i, start.Y})
	}
}

func (g Grid[T]) addRectangle(element T, start Vertex, end Vertex) {
	for i := start.X; i < end.X; i++ {
		for j := start.Y; j < end.Y; j++ {
			g.add(element, Vertex{i, j})
		}
	}
}
