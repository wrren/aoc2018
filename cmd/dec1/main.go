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
	changes, err := util.ReadIntegers(os.Args[2])
	util.Fatal(err)

	if os.Args[1] == "part1" {
		result, err := part1(changes)
		util.Fatal(err)
		fmt.Printf("Frequency: %d", result)
	} else if os.Args[1] == "part2" {
		result, err := part2(changes)
		util.Fatal(err)
		fmt.Printf("Repeated Frequency: %d", result)
	} else {
		util.Usage()
	}
}
