package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

type StepMap map[rune]*Step

type Step struct {
	ID           rune
	InProgress   bool
	IsFinished   bool
	Dependencies []*Step
}

func (s Step) TimeToComplete() int {
	return int(s.ID - 4)
}

func NextStep(m StepMap) (*Step, bool) {
	candidates := make([]*Step, 0)
	allFinished := true
	for _, v := range m {
		if v.IsFinished {
			continue
		}
		allFinished = false
		if v.InProgress {
			continue
		}
		isCandidate := true
		for _, s := range v.Dependencies {
			if s.IsFinished == false {
				isCandidate = false
				break
			}
		}
		if isCandidate {
			candidates = append(candidates, v)
		}
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].ID < candidates[j].ID
	})

	if len(candidates) == 0 {
		return nil, allFinished
	}
	return candidates[0], false
}

func ReadSteps(reader io.Reader) (StepMap, error) {
	buffered := bufio.NewReader(reader)
	m := StepMap{}
	for {
		s, err := buffered.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return m, err
		}
		err = NewStepFromString(s, m)
		if err != nil {
			return m, err
		}
	}
	return m, nil
}

func GetStep(id rune, m StepMap) *Step {
	_, ok := m[id]
	if ok == false {
		m[id] = &Step{
			ID:           id,
			IsFinished:   false,
			Dependencies: make([]*Step, 0),
		}
	}
	return m[id]
}

func NewStepFromString(s string, m StepMap) error {
	var id, dependency string
	_, err := fmt.Sscanf(s, "Step %s must be finished before step %s can begin.", &dependency, &id)

	if err != nil {
		return err
	}

	step := GetStep(rune(id[0]), m)
	dep := GetStep(rune(dependency[0]), m)
	step.Dependencies = append(step.Dependencies, dep)

	return nil
}
