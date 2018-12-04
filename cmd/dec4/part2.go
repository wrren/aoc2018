package main

import "fmt"

func part2(guards []Guard) (int, error) {
	if len(guards) == 0 {
		return 0, fmt.Errorf("Empty guard list provided")
	}

	id := 0
	max := 0
	minute := 0
	for _, g := range guards {
		for m, a := range g.SleepMap {
			if a > max {
				minute = m
				max = a
				id = g.ID
			}
		}
	}

	return minute * id, nil
}
