package main

import (
	"fmt"
)

type Square struct {
	X int
	Y int
}

type Claim struct {
	ID     int
	X      int
	Y      int
	Width  int
	Length int
}

type ClaimState int8

const (
	Unclaimed   ClaimState = 0
	SingleClaim ClaimState = 1
	MultiClaim  ClaimState = 2
)

type OverlayMap map[Square]ClaimState

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
			if m[square] == Unclaimed {
				m[square] = SingleClaim
			} else {
				m[square] = MultiClaim
			}
		}
	}
	return m
}

func (s Claim) OverlapsWith(m OverlayMap) bool {
	for i := 0; i < s.Width; i++ {
		for j := 0; j < s.Length; j++ {
			square := Square{i + s.X, j + s.Y}
			if m[square] > SingleClaim {
				return true
			}
		}
	}
	return false
}
