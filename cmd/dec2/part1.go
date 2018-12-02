package main

func checksum(id string, twos, threes *int) {
	characters := map[rune]int{}
	for _, r := range id {
		characters[r] = characters[r] + 1
	}
	var sawTwo, sawThree bool
	for _, count := range characters {
		if sawTwo && sawThree {
			return
		}
		if count == 2 && sawTwo == false {
			*twos = *twos + 1
			sawTwo = true
		}
		if count == 3 && sawThree == false {
			*threes = *threes + 1
			sawThree = true
		}
	}
}

func part1(ids []string) int {
	var twos, threes int
	for _, id := range ids {
		checksum(id, &twos, &threes)
	}
	return twos * threes
}
