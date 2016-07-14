package main

func calcLeftVoid(sp *span) (width int) {
	for ; width < len(sp.cells); width++ {
		if sp.cells[width] != 0 {
			break
		}
	}

	return
}

func calcRightVoid(sp *span) (width int) {
	for ; width < len(sp.cells); width++ {
		if sp.cells[len(sp.cells)-1-width] != 0 {
			break
		}
	}

	return
}

func calcInternalVoid(sp *span) (start, end int) {
	var j, max int
	for i := 0; i < len(sp.cells); i = j {
		cond := (sp.cells[i] == 0)
		for j = i + 1; j < len(sp.cells); j++ {
			if (sp.cells[j] == 0) != cond {
				break
			}
		}

		if cond {
			if w := j - i; w > max {
				max = w
				start, end = i, j
			}
		}
	}

	return
}
