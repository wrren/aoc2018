package main

import "io"

func part2(reader io.Reader) (int, error) {
	coordinates, err := ReadCoordinates(reader)
	if err != nil {
		return 0, err
	}
	box := MakeBoundingBox(coordinates)
	region := 0

	box.Iterate(func(x, y int) {
		total := 0
		for _, c := range coordinates {
			distance := c.ManhattanDistance(x, y)
			total = total + distance
		}

		if total < 10000 {
			region = region + 1
		}
	})

	return region, nil
}
