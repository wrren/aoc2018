package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaceMarble(t *testing.T) {
	game := NewMarbleGame()

	score := game.PlaceMarble(1)
	assert.Equal(t, 1, game.CurrentValue())
	assert.Equal(t, 0, score)

	score = game.PlaceMarble(2)
	assert.Equal(t, 2, game.CurrentValue())
	assert.Equal(t, 0, score)

	score = game.PlaceMarble(3)
	assert.Equal(t, []int{0, 2, 1, 3}, game.ToSlice())
	assert.Equal(t, 3, game.CurrentValue())
	assert.Equal(t, 0, score)

	score = game.PlaceMarble(4)
	assert.Equal(t, []int{0, 4, 2, 1, 3}, game.ToSlice())
	assert.Equal(t, 4, game.CurrentValue())
	assert.Equal(t, 0, score)

	score = game.PlaceMarble(5)
	assert.Equal(t, []int{0, 4, 2, 5, 1, 3}, game.ToSlice())
	assert.Equal(t, 5, game.CurrentValue())
	assert.Equal(t, 0, score)

	for i := 6; i < 23; i++ {
		score = game.PlaceMarble(i)
		assert.Equal(t, 0, score)
	}

	assert.Equal(t, 22, game.CurrentValue())
	score = game.PlaceMarble(23)
	assert.Equal(t, []int{0, 16, 8, 17, 4, 18, 19, 2, 20, 10, 21, 5, 22, 11, 1, 12, 6, 13, 3, 14, 7, 15}, game.ToSlice())
	assert.Equal(t, 19, game.CurrentValue())
	assert.Equal(t, 32, score)
}
