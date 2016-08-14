package main

import (
	"image"
	"image/gif"
	"os"
)

const frameDelay = 10

type gifAnimator struct {
	gif.GIF
}

func (g *gifAnimator) accumulate() {
	clone := image.NewPaletted(display.Rect, display.Palette)
	copy(clone.Pix, display.Pix)

	g.Image = append(g.Image, clone)
	g.Delay = append(g.Delay, frameDelay)
}

func (g *gifAnimator) encode(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	if err := gif.EncodeAll(f, &g.GIF); err != nil {
		f.Close()
		return err
	}

	return f.Close()
}
