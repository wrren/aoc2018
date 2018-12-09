package main

import (
	"io"
	"strings"
)

func part1(reader io.Reader) (string, error) {
	steps, err := ReadSteps(reader)
	sequence := strings.Builder{}
	if err != nil {
		return "", err
	}

	for {
		step, finished := NextStep(steps)
		if finished {
			break
		}
		step.IsFinished = true
		sequence.WriteRune(step.ID)
	}

	return sequence.String(), nil
}
