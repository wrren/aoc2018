package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) Adjacent() []Point {
	return []Point{
		Point{X: p.X, Y: p.Y - 1},
		Point{X: p.X, Y: p.Y + 1},
		Point{X: p.X - 1, Y: p.Y},
		Point{X: p.X + 1, Y: p.Y},
	}
}

func (p Point) ManhattanDistance(x, y int) int {
	return int(math.Abs(float64(p.X-x)) + math.Abs(float64(p.Y-y)))
}

func (p Point) String() string {
	return fmt.Sprintf("%d, %d", p.X, p.Y)
}
