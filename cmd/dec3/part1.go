package main

func part1(in []string) (int, error) {
	overlay := OverlayMap{}
	for _, s := range in {
		claim, err := NewClaimFromString(s)
		if err != nil {
			return 0, err
		}
		overlay = claim.Overlay(overlay)
	}

	overlap := 0
	for _, v := range overlay {
		if v > 1 {
			overlap = overlap + 1
		}
	}

	return overlap, nil
}
