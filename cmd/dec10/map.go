package main

import (
	"io"
	"math"
	"strings"
)

type LightPointMap struct {
	Points []LightPoint
	Map    map[Vector]bool
	Index  int
}

func ReadLightPointMap(reader io.Reader) (LightPointMap, error) {
	points, err := ReadLightPoints(reader)
	m := LightPointMap{
		Points: points,
		Map:    map[Vector]bool{},
	}
	for _, p := range points {
		m.Map[p.Position] = true
	}
	return m, err
}

func (m LightPointMap) VerticalLineCount() int {
	count := 0
	xs := map[Vector]bool{}
	for _, p := range m.Points {
		if xs[p.Position] {
			continue
		}
		xs[p.Position] = true
		length := 1
		vector := p.Position.Copy()
		for {
			vector.Y = vector.Y - 1
			if m.Map[vector] {
				length = length + 1
				xs[vector] = true
			} else {
				break
			}
		}
		vector = p.Position.Copy()
		for {
			vector.Y = vector.Y - 1
			if m.Map[vector] {
				length = length + 1
				xs[vector] = true
			} else {
				break
			}
		}
		if length > 6 {
			count = count + 1
		}
	}
	return count
}

func (m LightPointMap) BoundingBoxArea() int64 {
	left, right, top, bottom := math.MaxInt32, 0, math.MaxInt32, 0
	for _, p := range m.Points {
		if p.Position.X < left {
			left = p.Position.X
		}
		if p.Position.X > right {
			right = p.Position.X
		}
		if p.Position.Y < top {
			top = p.Position.Y
		}
		if p.Position.Y > bottom {
			bottom = p.Position.Y
		}
	}
	return int64(math.Abs(float64((right - left) * (bottom - top))))
}

func (m *LightPointMap) Advance() int64 {
	delete(m.Map, m.Points[m.Index].Position)
	m.Points[m.Index].Advance()
	m.Map[m.Points[m.Index].Position] = true
	m.Index = (m.Index + 1) % len(m.Points)
	return m.BoundingBoxArea()
}

func (m *LightPointMap) Retreat() int64 {
	m.Index = m.Index - 1
	if m.Index < 0 {
		m.Index = len(m.Points) - 1
	}
	delete(m.Map, m.Points[m.Index].Position)
	m.Points[m.Index].Retreat()
	m.Map[m.Points[m.Index].Position] = true
	return m.BoundingBoxArea()
}

func (m LightPointMap) Print() string {
	builder := strings.Builder{}
	left, right, top, bottom := math.MaxInt32, 0, math.MaxInt32, 0
	for _, p := range m.Points {
		if p.Position.X < left {
			left = p.Position.X
		}
		if p.Position.X > right {
			right = p.Position.X
		}
		if p.Position.Y < top {
			top = p.Position.Y
		}
		if p.Position.Y > bottom {
			bottom = p.Position.Y
		}
	}
	for y := 0; y < bottom-top; y++ {
		for x := 0; x < right-left; x++ {
			if m.Map[Vector{X: x + left, Y: y + top}] == true {
				builder.WriteRune('#')
			} else {
				builder.WriteRune('.')
			}
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}
