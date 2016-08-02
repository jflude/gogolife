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
			if a {
				if x < extent.Min.X {
					extent.Min.X = x
				}

				if x > extent.Max.X {
					extent.Max.X = x
				}
			}
		}

		x++
	}

	if y < extent.Min.Y {
		extent.Min.Y = y
	}

	if y > extent.Max.Y {
		extent.Max.Y = y
	}
}
