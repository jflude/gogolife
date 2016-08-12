package main

const patternDensity = 50

var patternSize int = displayWidth / 2

func sowPattern(kind int) (pat string) {
	switch kind {
	case 0:
		pat += randomPattern()
	case 1:
		pat += solidPattern()
	case 2:
		pat += agarPattern()
	}

	return
}
