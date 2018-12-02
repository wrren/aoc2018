package main

import (
	"fmt"
	"os"

	"github.com/wrren/aoc2018/internal/util"
)

func usage() {
	fmt.Println("usage:")
	fmt.Println("dec1 <input_file>")
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 {
		usage()
	}
	changes, err := util.ReadIntegers(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	seen := map[int]bool{0: true}

	var sum int
	for true {
		for _, c := range changes {
			sum = sum + c
			_, ok := seen[sum]
			if ok {
				fmt.Printf("Sum: %d", sum)
				os.Exit(0)
			}
			seen[sum] = true
		}
	}
}
