package main

import (
	"fmt"
	"os"

	"github.com/wrren/aoc2018/internal/util"
)

func main() {
	if len(os.Args) < 3 {
		util.Usage()
	}

	recordStrings, err := util.ReadStrings(os.Args[2])
	util.Fatal(err)

	records := make([]Record, 0, len(recordStrings))
	for _, r := range recordStrings {
		record, err := NewRecord(r)
		util.Fatal(err)
		records = append(records, record)
	}
	SortRecords(records)

	guards := GuardMap{}
	guard := NewGuard(0)
	started := false
	for _, r := range records {
		switch r.Event {
		case ShiftStarted:
			if started == false {
				started = true
				guard.ID = r.ID
			} else {
				guard.WokeUp(r.Time)
				guards[guard.ID] = guard
				_, ok := guards[r.ID]
				if ok == false {
					guard = NewGuard(r.ID)
				} else {
					guard = guards[r.ID]
				}
			}
		case FellAsleep:
			guard.FellAsleep(r.Time)
		case WokeUp:
			guard.WokeUp(r.Time)
		}
	}

	guardList := make([]Guard, 0, len(guards))
	for _, v := range guards {
		guardList = append(guardList, v)
	}

	if os.Args[1] == "part1" {
		result, err := part1(guardList)
		util.Fatal(err)
		fmt.Printf("Result: %d", result)
	} else if os.Args[1] == "part2" {
		result, err := part2(guardList)
		util.Fatal(err)
		fmt.Printf("Result: %d", result)
	} else {
		util.Usage()
	}
}
