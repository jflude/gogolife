package main

func pruneSpans(r *row) bool {
	for sp := r.spans.next; sp != &r.spans; sp = sp.next {
		if sp.near == 0 {
			removeSpan(sp)
		}
	}

	return r.spans.next == &r.spans
}
