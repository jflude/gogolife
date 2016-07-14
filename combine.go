package main

func combineSpans(r *row) {
	for sp := r.spans.next; sp != &r.spans; {
		if sp.right != sp.next.left {
			sp = sp.next
			continue
		}

		sp.cells = append(sp.cells, sp.next.cells...)
		sp.right = sp.next.right
		sp.near += sp.next.near

		for i := range r.changes {
			if r.changes[i].sp == sp.next {
				r.changes[i].sp = sp
			}
		}

		removeSpan(sp.next)
	}
}
