package main

import (
	"container/ring"
)

type MarbleGame struct {
	Current *ring.Ring
	Marbles *ring.Ring
}

func NewMarbleGame() MarbleGame {
	game := MarbleGame{Marbles: ring.New(1)}
	game.Marbles.Value = 0
	game.Current = game.Marbles
	return game
}

func (m MarbleGame) CurrentValue() int {
	return m.Current.Value.(int)
}

func (m MarbleGame) ToSlice() []int {
	slice := make([]int, 1, m.Marbles.Len())
	slice[0] = m.Marbles.Value.(int)

	current := m.Marbles.Next()
	for current != m.Marbles {
		slice = append(slice, current.Value.(int))
		current = current.Next()
	}
	return slice
}

func (m *MarbleGame) PlaceMarble(marble int) int {
	if marble%23 != 0 {
		m.Current = m.Current.Next()
		new := ring.New(1)
		new.Value = marble
		m.Current.Link(new)
		m.Current = m.Current.Next()
		return 0
	}

	score := marble
	for i := 0; i < 8; i++ {
		m.Current = m.Current.Prev()
	}
	score = score + m.Current.Next().Value.(int)
	m.Current.Unlink(1)
	m.Current = m.Current.Next()

	return score
}
