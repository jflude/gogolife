package main

import (
	"math/rand"
	"strings"
)

func agarPattern(size int) string {
	third := size / 3
	contam := third/3 + rand.Intn(third/3)

	pat := "0 0\n"
	for i := 0; i < third; i++ {
		var s string
		for j := 0; j < third; j++ {
			s += "OO."
		}

		s += "\n" + s + "\n"
		if i == contam {
			n := third + rand.Intn(third)
			if n%3 == 2 {
				n--
			}

			s += strings.Repeat(".", n) + "O"
		}

		pat += ".\n" + s
	}

	return pat
}
