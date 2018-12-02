package main

import (
	"fmt"
	"os"

	"github.com/wrren/aoc2018/internal/util"
)

func usage() {
	fmt.Println("usage:")
	fmt.Println("dec1 <part1|part2> <input_file>")
	os.Exit(1)
}

func fatal(err error) {
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}
	changes, err := util.ReadIntegers(os.Args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if os.Args[1] == "part1" {
		result, err := part1(changes)
		fatal(err)
		fmt.Printf("Frequency: %d", result)
	} else if os.Args[1] == "part2" {
		result, err := part2(changes)
		fatal(err)
		fmt.Printf("Repeated Frequency: %d", result)
	} else {
		usage()
	}
}
