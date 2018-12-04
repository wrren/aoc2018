package main

import (
	"fmt"
	"sort"
)

func part1(guards []Guard) (int, error) {
	if len(guards) == 0 {
		return 0, fmt.Errorf("Empty guard list provided")
	}
	sort.Slice(guards, func(i, j int) bool {
		return guards[j].AsleepFor < guards[i].AsleepFor
	})

	guard := guards[0]
	max := 0
	minute := 0
	for m, a := range guard.SleepMap {
		if a > max {
			minute = m
			max = a
		}
	}

	return minute * guard.ID, nil
}
