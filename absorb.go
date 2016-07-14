package main

func isBehindRow(r *row, y int) bool {
	return y > r.prev.y && y < r.y
}

func isBehindSpan(sp *span, x int) bool {
	return x >= sp.prev.right && x < sp.left
}

func isWithinSpan(sp *span, x int) bool {
	return x >= sp.left && x < sp.right
}

func absorbCellRight(sp *span, x int) bool {
	if x-cellChunk >= sp.right {
		return false
	}

	sp.cells = append(sp.cells, make([]int8, cellChunk)...)
	sp.right += cellChunk
	return true
}

func absorbCellLeft(sp *span, x int) bool {
	if sp.left >= x+cellChunk {
		return false
	}

	sp.cells = append(make([]int8, cellChunk), sp.cells...)
	sp.left -= cellChunk
	return true
}
