package main

import (
	"fmt"
	"time"
)

var generation, population int64

func generate(maxGen int64) bool {
	now := time.Now()
	population = 0

	startJob(&apply)
	startJob(&calc)
	startJob(&tidy)

	generation++

	fmt.Println("### GENERATION", generation, population,
		float64(time.Now().Sub(now).Nanoseconds())/1000, bounds)

	printWorld()
	return (maxGen == 0 || generation < maxGen) && rows.next != &rows
}
