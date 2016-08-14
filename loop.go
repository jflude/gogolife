package main

import (
	"fmt"
	"os"
	"time"
)

func loop(max int) error {
	var ani gifAnimator
	for {
		if generation%*period == 0 {
			ani.accumulate()
		}

		if generation >= max {
			break
		}

		var g bool
		if *debug > 0 {
			now := time.Now()
			g = generate()

			fmt.Fprintln(os.Stderr, "### GENERATION", generation, population,
				float64(time.Now().Sub(now).Nanoseconds())/1000, extent)

			if *debug > 1 {
				printWorld()
			}
		} else {
			g = generate()
		}

		if !g {
			break
		}
	}

	return ani.encode("life")
}
