package main

import "fmt"

type Square struct {
	X int
	Y int
}

type OverlayMap map[Square]int

type Claim struct {
	ID     int
	X      int
	Y      int
	Width  int
	Length int
}

func NewClaimFromString(descriptor string) (Claim, error) {
	s := Claim{}
	_, err := fmt.Sscanf(descriptor, "#%d @ %d,%d: %dx%d", &(s.ID), &(s.X), &(s.Y), &(s.Width), &(s.Length))
	if err != nil {
		return s, err
	}
	return s, nil
}

func (s Claim) Overlay(m OverlayMap) OverlayMap {
	for i := 0; i < s.Width; i++ {
		for j := 0; j < s.Length; j++ {
			square := Square{i + s.X, j + s.Y}
			m[square] = m[square] + 1
		}
	}
	return m
}

func (s Claim) OverlapsWith(m OverlayMap) bool {
	for i := 0; i < s.Width; i++ {
		for j := 0; j < s.Length; j++ {
			square := Square{i + s.X, j + s.Y}
			if m[square] > 1 {
				return true
			}
		}
	}
	return false
}
