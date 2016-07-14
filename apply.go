package main

type applyJob struct {
	expanded [][]*row
}

var apply applyJob

func (j *applyJob) start() {
	j.expanded = make([][]*row, stride)
}

func (j *applyJob) work(id int, r *row) {
	for _, chg := range r.changes {
		cell := findCell(chg.sp, chg.x)
		changeState(cell)
		drawCell(chg.x, r.y, *cell)

		var delta int8
		if isAlive(*cell) {
			r.pop++
			delta = 1 << 1
		} else {
			r.pop--
			delta = -(1 << 1)
		}

		r.effects = affectNear(r.effects, chg.x, delta, delta, delta)
	}

	r.changes = r.changes[:0]

	if len(r.effects) > 0 {
		j.expanded[id] = append(j.expanded[id], r)
	}
}

func (j *applyJob) finish() {
	for _, w := range j.expanded {
		for _, r := range w {
			if r.prev.y != r.y-1 {
				insertBeforeRow(newRow(r.y-1), r)
			}

			if r.next.y != r.y+1 {
				insertAfterRow(newRow(r.y+1), r)
			}
		}
	}
}
