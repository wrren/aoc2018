package main

import (
	"fmt"
	"os"

	"github.com/wrren/aoc2018/internal/util"
)

func usage() {
	fmt.Println("usage:")
	fmt.Println("dec2 <part1|part2> <input_file>")
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

	claims, err := util.ReadStrings(os.Args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if os.Args[1] == "part1" {
		overlap, err := part1(claims)
		fatal(err)
		fmt.Printf("Overlapping Squares: %d", overlap)
	} else if os.Args[1] == "part2" {
		claim, err := part2(claims)
		fatal(err)
		fmt.Printf("Claim ID: %d", claim.ID)
	} else {
		usage()
	}
}
