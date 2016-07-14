package main

func transferChanges(r *row, src, dest *span) {
	for i := range r.changes {
		if r.changes[i].sp == src {
			for j := range r.changes[i:] {
				if r.changes[i+j].sp != src {
					break
				}

				if isWithinSpan(dest, r.changes[i+j].x) {
					r.changes[i+j].sp = dest
				}
			}

			break
		}
	}
}
