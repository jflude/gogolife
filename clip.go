package main

const minClipSize = 8

func clipSpans(r *row) {
	for sp := r.spans.next; sp != &r.spans; sp = sp.next {
		if w := calcLeftVoid(sp); w >= minClipSize {
			if len(r.changes) > 0 {
				max := r.changes[0].x - sp.left - 1
				if w > max {
					w = max
				}
			}

			if w = (w &^ (cellChunk - 1)) - cellChunk; w > 0 {
				sp.cells = append([]int8(nil), sp.cells[w:]...)
				sp.left += w
			}
		}

		if w := calcRightVoid(sp); w >= minClipSize {
			if len(r.changes) > 0 {
				max := sp.right - r.changes[len(r.changes)-1].x - 1
				if w > max {
					w = max
				}
			}

			if w = (w &^ (cellChunk - 1)) - cellChunk; w > 0 {
				sp.cells = append([]int8(nil), sp.cells[:len(sp.cells)-w]...)
				sp.right -= w
			}
		}
	}
}
