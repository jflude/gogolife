package main

import (
	"math/rand"
	"strings"
)

func agarPattern() (pat string) {
	contam := rand.Intn(patternSize / 3)
	for i := 0; i < patternSize/3; i++ {
		var s string
		for j := 0; j < patternSize/3; j++ {
			s += "OO."
		}

		s += "\n" + s + "\n"
		if i == contam {
			n := rand.Intn(patternSize - 1)
			if n%3 == 2 {
				n--
			}

			s += strings.Repeat(".", n) + "O"
		}

		pat += ".\n" + s
	}

	return
}
