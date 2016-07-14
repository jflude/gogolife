package main

import "sort"

type placeJob struct{}

var place placeJob

func (j placeJob) start() {
	// empty
}

func (j placeJob) finish() {
	// empty
}

func (j placeJob) work(id int, r *row) {
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

func placeCells(x, y int, alive []bool) {
	r := findRow(nil, y)
	var sp *span

	for _, a := range alive {
		sp = findSpan(r, sp, x)
		if a != isAlive(*findCell(sp, x)) {
			r.changes = append(r.changes, change{sp, x})
		}

		x++
	}
}

func placeReady() {
	startJob(&place)
}
