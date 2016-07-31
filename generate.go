package main

import (
	"fmt"
	"time"
)

var generation, population int64

func generate() bool {
	now := time.Now()
	population = 0

	startJob(&apply)
	startJob(&calc)
	startJob(&tidy)

	generation++

	fmt.Println("### GENERATION", generation, population,
		float64(time.Now().Sub(now).Nanoseconds())/1000, extent)

	printWorld()
	return rows.next != &rows
}
