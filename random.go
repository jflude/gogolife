package main

import "math/rand"

const density = 50

func randomPattern(size int) string {
	pat := "0 0\n"
	for i := 0; i < size; i++ {
		var s string
		for j := 0; j < size; j++ {
			if rand.Intn(100) < density {
				s += "O"
			} else {
				s += "."
			}
		}

		pat += s + "\n"
	}

	return pat
}
