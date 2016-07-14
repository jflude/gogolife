package main

func affectNear(eff []effect, orig int, delta ...int8) []effect {
	if len(eff) == 0 {
		for i := 0; i < 3; i++ {
			x := orig - 1 + i
			eff = append(eff, effect{x, delta[i]})
		}
	} else {
		for i := 0; i < 3; i++ {
			x := orig - 1 + i
			if x == eff[len(eff)-2].x {
				eff[len(eff)-2].delta += delta[i]
			} else if x == eff[len(eff)-1].x {
				eff[len(eff)-1].delta += delta[i]
			} else {
				eff = append(eff, effect{x, delta[i]})
			}
		}
	}

	return eff
}

func calcNear(sp *span) {
	sp.near = 0
	for _, cell := range sp.cells {
		sp.near += int(neighbourCount(cell))
	}
}
