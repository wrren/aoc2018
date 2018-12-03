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

	claims, err := util.ReadStrings(os.Args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if os.Args[1] == "part1" {
		overlap, err := part1(claims)
		util.Fatal(err)
		fmt.Printf("Overlapping Squares: %d", overlap)
	} else if os.Args[1] == "part2" {
		claim, err := part2(claims)
		util.Fatal(err)
		fmt.Printf("Claim ID: %d", claim.ID)
	} else {
		util.Usage()
	}
}
