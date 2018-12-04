package main

import (
	"fmt"
	"sort"
	"time"
)

type EventType int8

const (
	ShiftStarted EventType = 0
	FellAsleep   EventType = 1
	WokeUp       EventType = 2
)

type Record struct {
	Time  time.Time
	ID    int
	Event EventType
}

func NewRecord(event string) (Record, error) {
	r := Record{}
	runes := []rune(event)
	when := string(runes[1 : len("2006-01-02 15:04")+1])
	content := string(runes[len("[2006-01-02 15:04] "):])

	t, err := time.Parse("2006-01-02 15:04", when)
	if err != nil {
		return r, fmt.Errorf("Failed to extract timestamp from record %s: %s", when, err.Error())
	}
	r.Time = t

	if content == "wakes up" {
		r.Event = WokeUp
	} else if content == "falls asleep" {
		r.Event = FellAsleep
	} else {
		_, err := fmt.Sscanf(content, "Guard #%d begins shift", &(r.ID))
		if err != nil {
			return r, fmt.Errorf("Record content does not match expected format")
		}
	}
	return r, nil
}

func SortRecords(records []Record) {
	sort.Slice(records, func(i, j int) bool {
		return records[i].Time.Before(records[j].Time)
	})
}
