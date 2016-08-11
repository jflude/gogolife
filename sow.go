package main

import (
	"fmt"
	"math/rand"
	"time"
)

const patternDensity = 50

var patternSize int = displayWidth / 2

func sowPattern(kind int) (pat string) {
	seed := time.Now().UnixNano()
	fmt.Println("Seed is", seed)
	rand.Seed(seed)

	pat = "0 0\n"

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
