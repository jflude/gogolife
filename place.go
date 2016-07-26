package main

import "image"

var extent image.Rectangle

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

	extent = extent.Union(image.Rect(x, y, x, y))
}
