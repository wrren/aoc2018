package main

import (
	"fmt"
)

func part2(in []string) (Claim, error) {
	overlay := OverlayMap{}
	claims := make([]Claim, 0, len(in))
	for _, s := range in {
		claim, err := NewClaimFromString(s)
		if err != nil {
			return Claim{}, err
		}
		overlay = claim.Overlay(overlay)
		claims = append(claims, claim)
	}

	for _, c := range claims {
		if c.OverlapsWith(overlay) == false {
			return c, nil
		}
	}

	return Claim{}, fmt.Errorf("No claim found without overlap")
}
