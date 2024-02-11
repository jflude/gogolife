package main

import (
	"fmt"
	"os"
	"time"
)

func loop(ani animator, filename string, maxGen int) error {
	var done bool
	for {
		if generation%*skip == 0 {
			ani.accumulate()
		}

		if done || generation >= maxGen {
			break
		}

		if *debug < 2 {
			done = !generate()
		} else {
			begin := time.Now()
			done = !generate()

			fmt.Fprintln(os.Stderr, "### GENERATION",
				generation, population,
				float64(time.Since(begin).Nanoseconds())/1000,
				extent)

			if *debug > 2 {
				printWorld()
			}
		}
	}

	return ani.encode(filename)
}
