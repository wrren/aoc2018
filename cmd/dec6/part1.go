package main

import (
	"io"
	"math"
)

func part1(reader io.Reader) (int, error) {
	coordinates, err := ReadCoordinates(reader)
	if err != nil {
		return 0, err
	}

	box := MakeBoundingBox(coordinates)
	for i, c := range coordinates {
		coordinates[i].IsBorder = box.Intersects(c.Point)
	}

	m := map[rune]int{}
	max := 0
	box.Iterate(func(x, y int) {
		min := math.MaxInt32
		id := SharedSpace
		var target *Coordinate
		for _, c := range coordinates {
			distance := c.ManhattanDistance(x, y)
			if distance == min {
				id = SharedSpace
			} else if distance < min {
				target = &c
				min = distance
				id = c.ID
			}
		}

		if id != SharedSpace && target.IsBorder == false {
			m[id] = m[id] + 1
			if m[id] > max {
				max = m[id]
			}
		}
	})

	return max, nil
}
