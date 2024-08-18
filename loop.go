package main

import (
	"fmt"
	"image"
	"os"
	"time"
)

func loop() error {
	maxGen := *gens * *skip
	var done bool
	for {
		if generation%*skip == 0 {
			img := image.NewPaletted(display.Rect, display.Palette)
			copy(img.Pix, display.Pix)
			if *caption {
				cap := fmt.Sprint("gen %v", generation)
				_ = cap // TODO: draw cap on img
			}
			ani.accumulate(img, frameDelay)
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
	f, err := os.Create(*file)
	if err != nil {
		return err
	}
	if err = ani.encode(f); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}
