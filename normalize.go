package main

import "sort"

type normalizeJob struct{}

var normalize normalizeJob

func (j normalizeJob) start() {
	// empty
}

func (j normalizeJob) work(id int, r *row) {
	if len(r.changes) == 0 {
		return
	}

	sort.Sort(ChangeSlice(r.changes))

	k := 0
	for i := 1; i < len(r.changes); i++ {
		if r.changes[i].x != r.changes[k].x {
			k++
			r.changes[k] = r.changes[i]
		}
	}

	r.changes = r.changes[:k+1]
}

func (j normalizeJob) finish() {
	// empty
}
