package main

func part1(changes []int) (int, error) {
	sum := 0
	for _, v := range changes {
		sum = sum + v
	}
	return sum, nil
}
