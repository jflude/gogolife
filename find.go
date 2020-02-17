package main

func findCell(sp *span, x int) *int8 {
	return &sp.cells[x-sp.left]
}

func findSpan(r *row, from *span, x int) *span {
	if from == nil {
		from = r.spans.next
	}

	for {
		if from == &r.spans {
			if absorbCellRight(from.prev, x) {
				return from.prev
			}

			return insertAfterSpan(newSpan(x), r.spans.prev)
		}

		if isWithinSpan(from, x) {
			return from
		}

		if isBehindSpan(from, x) {
			if absorbCellRight(from.prev, x) {
				return from.prev
			} else if absorbCellLeft(from, x) {
				return from
			}

			return insertBeforeSpan(newSpan(x), from)
		}

		from = from.next
	}
}

func findRow(from *row, y int) *row {
	if from == nil {
		from = rows.next
	}

	for {
		if from == &rows {
			return insertAfterRow(newRow(y), rows.prev)
		}

		if from.y == y {
			return from
		}

		if isBehindRow(from, y) {
			return insertBeforeRow(newRow(y), from)
		}

		from = from.next
	}
}
