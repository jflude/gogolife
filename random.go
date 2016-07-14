package main

import "math/rand"

func randomPattern() (pat string) {
	for i := 0; i < patternSize; i++ {
		var s string
		for j := 0; j < patternSize; j++ {
			if rand.Intn(100) < patternDensity {
				s += "O"
			} else {
				s += "."
			}
		}

		pat += s + "\n"
	}

	return
}
