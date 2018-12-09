package main

import (
	"bufio"
	"fmt"
	"io"
)

const SharedSpace = '.'

type Coordinate struct {
	Point
	ID       rune
	Area     int
	IsBorder bool
}

func ReadCoordinates(reader io.Reader) ([]Coordinate, error) {
	buffered := bufio.NewReader(reader)
	coordinates := make([]Coordinate, 0)
	id := 'A'
	for {
		v, err := buffered.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return coordinates, nil
			}
			return coordinates, err
		}
		c, err := NewCoordinateFromString(v)
		c.ID = rune(id)
		coordinates = append(coordinates, c)
		id++
	}
	return coordinates, nil
}

func NewCoordinateFromString(s string) (Coordinate, error) {
	c := Coordinate{Area: 1}
	_, err := fmt.Sscanf(s, "%d, %d", &(c.X), &(c.Y))
	return c, err
}
