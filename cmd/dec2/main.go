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

	ids, err := util.ReadStrings(os.Args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if os.Args[1] == "part1" {
		fmt.Printf("Checksum: %d", part1(ids))
	} else if os.Args[1] == "part2" {
		result, err := part2(ids)
		util.Fatal(err)
		fmt.Printf("Common Letters: %s", result)
	} else {
		util.Usage()
	}
}
