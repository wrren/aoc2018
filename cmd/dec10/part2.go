package main

import (
	"io"
	"math"
)

func part2(reader io.Reader) (int, error) {
	m, err := ReadLightPointMap(reader)
	seconds := 0
	if err != nil {
		return 0, err
	}
	min := int64(math.MaxInt64)
	for {
		area := m.Advance()
		seconds = seconds + 1
		if area < min {
			min = area
		} else if area > min {
			m.Print()
			break
		}
	}
	return (seconds / len(m.Points)), nil
}
