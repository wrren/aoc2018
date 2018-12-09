package main

import (
	"fmt"
	"math"
)

type BoundingBox struct {
	Top, Bottom, Left, Right int
}

func (b BoundingBox) String() string {
	return fmt.Sprintf("%d, %d, %d, %d", b.Top, b.Bottom, b.Left, b.Right)
}

func (b BoundingBox) Contains(p Point) bool {
	return p.X >= b.Left && p.X <= b.Right && p.Y >= b.Top && p.Y <= b.Bottom
}

func (b BoundingBox) Intersects(p Point) bool {
	return p.X == b.Left || p.X == b.Right || p.Y == b.Top || p.Y == b.Bottom
}

func MakeBoundingBox(coordinates []Coordinate) BoundingBox {
	b := BoundingBox{Left: math.MaxInt32, Right: 0, Top: math.MaxInt32, Bottom: 0}

	for _, c := range coordinates {
		if c.X < b.Left {
			b.Left = c.X
		}
		if c.X > b.Right {
			b.Right = c.X
		}
		if c.Y < b.Top {
			b.Top = c.Y
		}
		if c.Y > b.Bottom {
			b.Bottom = c.Y
		}
	}
	return b
}

func (b BoundingBox) Iterate(f func(x, y int)) {
	for i := b.Top; i <= b.Bottom; i++ {
		for j := b.Left; j <= b.Right; j++ {
			f(j, i)
		}
	}
}
