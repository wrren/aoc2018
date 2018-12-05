package main

import "testing"

func TestPolymerReaction(t *testing.T) {
	var units = [...]rune{'d', 'a', 'b', 'A', 'c', 'C', 'a', 'C', 'B', 'A', 'c', 'C', 'c', 'a', 'D', 'A'}
	polymer := NewPolymer()
	for _, u := range units {
		polymer.Chain(u)
	}
	if len(polymer.Buffer) != 10 {
		t.Fatalf("Polymer length != 10. Actual length: %d", len(polymer.Buffer))
	}
}
