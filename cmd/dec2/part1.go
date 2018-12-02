package main

type counter map[int]int

func checksum(id string, counter *counter) {
	characters := map[rune]int{}
	for _, r := range id {
		characters[r] = characters[r] + 1
	}
	for k, _ := range *counter {
		for _, count := range characters {
			if count == k {
				(*counter)[k] = (*counter)[k] + 1
				break
			}
		}
	}
}

func part1(ids []string) int {
	counter := counter{
		2: 0,
		3: 0,
	}
	for _, id := range ids {
		checksum(id, &counter)
	}
	return counter[2] * counter[3]
}
