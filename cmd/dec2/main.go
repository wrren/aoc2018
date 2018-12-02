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

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	ids, err := util.ReadStrings(os.Args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if os.Args[1] == "part1" {
		fmt.Printf("Checksum: %d", part1(ids))
	} else if os.Args[1] == "part2" {
		fmt.Printf("Common Letters: %s", part2(ids))
	} else {
		usage()
	}
}
