package main

import (
	"time"
)

type GuardMap map[int]Guard

type Guard struct {
	ID           int
	SleepMap     map[int]int
	IsAsleep     bool
	AsleepFor    int
	FellAsleepOn int
}

func NewGuard(id int) Guard {
	return Guard{ID: id, SleepMap: map[int]int{}}
}

func (g *Guard) FellAsleep(time time.Time) {
	g.FellAsleepOn = time.Minute()
	g.IsAsleep = true
}

func (g *Guard) WokeUp(time time.Time) {
	if g.IsAsleep {
		for i := g.FellAsleepOn; i < time.Minute(); i++ {
			g.SleepMap[i] = g.SleepMap[i] + 1
			g.AsleepFor = g.AsleepFor + 1
		}
		g.IsAsleep = false
	}
}
