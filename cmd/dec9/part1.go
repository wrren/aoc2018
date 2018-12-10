package main

import (
	"bufio"
	"fmt"
	"io"
)

func part1(reader io.Reader) (int, error) {
	buffered := bufio.NewReader(reader)
	description, err := buffered.ReadString('\n')
	if err != nil {
		return 0, err
	}
	var players, points int
	_, err = fmt.Sscanf(description, "%d players; last marble is worth %d points", &players, &points)
	if err != nil {
		return 0, err
	}

	game := NewMarbleGame()
	scores := make([]int, players)
	max := 0
	player := 0
	for i := 1; i <= points; i++ {
		scores[player] = scores[player] + game.PlaceMarble(i)
		if scores[player] > max {
			max = scores[player]
		}
		player = (player + 1) % len(scores)
	}

	return max, nil
}
