package main

import (
	"io"
)

func part2(reader io.Reader) (int, error) {
	steps, err := ReadSteps(reader)

	if err != nil {
		return 0, err
	}
	time := 0
	workers := make([]Worker, 5)

	for {
		step, finished := NextStep(steps)
		if finished {
			break
		}

		worker := workers[0]
		if step != nil && worker.TimeLeft == 0 {
			workers[0].StartWorkOn(step)
			SortWorkers(workers)
		} else {
			next := NextWorkerToFinish(workers)
			left := next.TimeLeft
			time = time + left
			for i := range workers {
				workers[i].WorkFor(left)
			}
			SortWorkers(workers)
		}
	}

	return time, nil
}
