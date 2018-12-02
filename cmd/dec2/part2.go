package main

import (
	"math"
	"sort"
	"strings"
	"unicode/utf8"
)

type ID struct {
	id       string
	checksum int
}

func NewID(id string) ID {
	result := 0
	for _, r := range id {
		result = result + int(r)
	}
	return ID{
		id:       id,
		checksum: result,
	}
}

func (id ID) Match(other ID) (bool, bool) {
	if math.Abs(float64(id.checksum-other.checksum)) < 26.0 {
		differences := 0
		for i, c := range id.id {
			if differences > 1 {
				return true, false
			}
			other, _ := utf8.DecodeRuneInString(other.id[i:])
			if c != other {
				differences = differences + 1
			}
		}
		return true, true
	}
	return false, false
}

func (id ID) Common(other ID) string {
	builder := strings.Builder{}
	for i, c := range id.id {
		other, _ := utf8.DecodeRuneInString(other.id[i:])
		if c == other {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func part2(idStrings []string) string {
	ids := make([]ID, 0, len(idStrings))
	for _, id := range idStrings {
		ids = append(ids, NewID(id))
	}
	sort.Slice(ids, func(i, j int) bool {
		return ids[i].checksum < ids[j].checksum
	})

	for i := range ids {
		for j := i + 1; j < len(ids); j++ {
			quick, deep := ids[i].Match(ids[j])
			if quick && deep {
				return ids[i].Common(ids[j])
			} else if !quick {
				break
			}
		}
	}
	return ""
}
