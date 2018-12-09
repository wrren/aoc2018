package main

import (
	"math"
	"sort"
)

type Worker struct {
	TimeLeft  int
	WorkingOn *Step
	Total     int
}

func NextWorkerToFinish(w []Worker) *Worker {
	for i := range w {
		if w[i].TimeLeft != 0 {
			return &(w[i])
		}
	}
	return nil
}

func SortWorkers(w []Worker) {
	sort.Slice(w, func(i, j int) bool {
		return w[i].TimeLeft < w[j].TimeLeft
	})
}

func (w *Worker) StartWorkOn(s *Step) {
	w.TimeLeft = s.TimeToComplete()
	w.WorkingOn = s
	w.WorkingOn.InProgress = true
}

func (w *Worker) WorkFor(time int) {
	w.Total = w.Total + time
	w.TimeLeft = int(math.Max(float64(w.TimeLeft-time), 0))
	if w.TimeLeft == 0 && w.WorkingOn != nil {
		w.WorkingOn.IsFinished = true
		w.WorkingOn = nil
	}
}
