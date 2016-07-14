package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	patternSize    = 20
	patternDensity = 50
)

func sowPattern() (pat string) {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	fmt.Println("Seed is", seed)

	pat = fmt.Sprintf("%d 0\n", patternSize/2)
	//pat += randomPattern()
	//pat += solidPattern()
	pat += agarPattern()
	return
}
