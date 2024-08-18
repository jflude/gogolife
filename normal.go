package main

import "sort"

type normalJob struct{}

var normal normalJob

func (j normalJob) start() {
	origin.X = (displayWidth - extent.Dx()) / 2
	origin.Y = (displayHeight - extent.Dy()) / 2
}

func (j normalJob) work(id int, r *row) {
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

func (j normalJob) finish() {
	// empty
}
