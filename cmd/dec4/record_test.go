package main

import "testing"

var events = [...]string{
	"[1518-10-23 23:49] Guard #659 begins shift",
	"[1518-08-06 00:07] falls asleep",
	"[1518-06-24 00:58] wakes up",
	"[1518-04-27 00:13] falls asleep",
	"[1518-06-08 00:01] Guard #179 begins shift",
	"[1518-06-03 00:00] Guard #659 begins shift",
}

func TestNewRecord(t *testing.T) {
	record, err := NewRecord(events[0])
	if err != nil {
		t.Fatal(err)
	}
	if record.ID != 659 {
		t.Fatal("Guard ID Incorrect")
	}
	if record.Event != ShiftStarted {
		t.Fatal("Incorrect Event Type")
	}
}
