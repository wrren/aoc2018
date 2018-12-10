package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type LightPoint struct {
	Position Vector
	Velocity Vector
}

func ReadLightPoints(reader io.Reader) ([]LightPoint, error) {
	points := make([]LightPoint, 0)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		s := scanner.Text()
		point, err := NewLightPointFromString(s)
		if err != nil {
			return points, err
		}
		points = append(points, point)
	}
	return points, scanner.Err()
}

func NewLightPointFromString(s string) (LightPoint, error) {
	point := LightPoint{}
	_, err := fmt.Sscanf(strings.Replace(s, " ", "", -1), "position=<%d,%d>velocity=<%d,%d>", &(point.Position.X), &(point.Position.Y), &(point.Velocity.X), &(point.Velocity.Y))
	return point, err
}

func (p *LightPoint) Advance() {
	p.Position.X = p.Position.X + p.Velocity.X
	p.Position.Y = p.Position.Y + p.Velocity.Y
}

func (p *LightPoint) Retreat() {
	p.Position.X = p.Position.X - p.Velocity.X
	p.Position.Y = p.Position.Y - p.Velocity.Y
}
