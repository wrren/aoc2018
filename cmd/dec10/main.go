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

	reader, err := os.Open(os.Args[2])
	util.Fatal(err)

	if os.Args[1] == "part1" {
		message, err := part1(reader)
		util.Fatal(err)
		fmt.Print(message)
	} else if os.Args[1] == "part2" {
		seconds, err := part2(reader)
		util.Fatal(err)
		fmt.Printf("Seconds: %d\n", seconds)
	} else {
		util.Usage()
	}
}
