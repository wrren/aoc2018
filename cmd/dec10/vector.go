package main

type Vector struct {
	X, Y int
}

func (v Vector) Copy() Vector {
	return Vector{X: v.X, Y: v.Y}
}
