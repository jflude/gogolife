package main

type calcJob struct{}

var calc calcJob

func (j calcJob) start() {
	// empty
}

func (j calcJob) finish() {
	// empty
}

func (j calcJob) work(id int, r *row) {
	m0 := newMergeable(r.prev.effects)
	m1 := newMergeable(r.effects)
	m2 := newMergeable(r.next.effects)

	var sp *span
	for {
		aff, done := merge3Effects(m0, m1, m2)
		if done {
			break
		}

		sp = findSpan(r, sp, aff.x)
		sp.near += int(aff.delta)

		cell := findCell(sp, aff.x)
		*cell += aff.delta

		if isChanging(*cell) {
			r.changes = append(r.changes, change{sp, aff.x})
		}
	}
}
