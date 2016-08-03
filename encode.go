package main

import (
	"image"
	"image/gif"
	"os"
)

const frameDelay = 10

var animated gif.GIF

func accumulate() {
	clone := image.NewPaletted(display.Rect, display.Palette)
	copy(clone.Pix, display.Pix)

	animated.Image = append(animated.Image, clone)
	animated.Delay = append(animated.Delay, frameDelay)
}

func encode(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	if err := gif.EncodeAll(f, &animated); err != nil {
		f.Close()
		return err
	}

	return f.Close()
}
