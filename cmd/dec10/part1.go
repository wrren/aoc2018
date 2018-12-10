package main

import (
	"io"
	"math"
)

func part1(reader io.Reader) (string, error) {
	m, err := ReadLightPointMap(reader)
	message := ""
	if err != nil {
		return "", err
	}
	min := int64(math.MaxInt64)
	for {
		area := m.Advance()
		if area < min {
			min = area
		} else if area > min {
			message = m.Print()
			break
		}
	}
	return message, nil
}
