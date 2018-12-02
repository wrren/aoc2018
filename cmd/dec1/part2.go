package main

import "fmt"

func part2(changes []int) (int, error) {
	seen := map[int]bool{0: true}

	var sum int
	for true {
		for _, c := range changes {
			sum = sum + c
			_, ok := seen[sum]
			if ok {
				return sum, nil
			}
			seen[sum] = true
		}
	}
	return 0, fmt.Errorf("Failed to find matching frequency")
}
