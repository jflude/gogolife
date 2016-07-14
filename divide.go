package main

const (
	minSplitLength = 16
	minSplitSize   = 8
)

func divideSpans(r *row) {
	for sp := r.spans.next; sp != &r.spans; sp = sp.next {
		if len(sp.cells) < minSplitLength {
			continue
		}

		startMax, endMax := calcInternalVoid(sp)

		startMax = startMax&^(cellChunk-1) + cellChunk
		endMax = endMax&^(cellChunk-1) - cellChunk

		if endMax-startMax < minSplitSize {
			continue
		}

		spNew := insertAfterSpan(newSpan(sp.left+endMax), sp)

		setCells(spNew, append([]int8(nil), sp.cells[endMax:]...))
		setCells(sp, append([]int8(nil), sp.cells[:startMax]...))

		transferChanges(r, sp, spNew)
	}
}

func setCells(sp *span, cells []int8) {
	sp.cells = cells
	sp.right = sp.left + len(sp.cells)
	calcNear(sp)
}
