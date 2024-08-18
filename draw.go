package main

import (
	"image"
	"image/color"
	"io"
)

const (
	displayWidth  = 333 // in pixels
	displayHeight = 333
	frameDelay    = 100 // in milliseconds
)

type animator interface {
	accumulate(img *image.Paletted, delayMillis int)
	encode(w io.Writer) error
}

var display = image.NewPaletted(
	image.Rect(0, 0, 3*displayWidth, 3*displayHeight),
	color.Palette{color.White, color.Black})

var origin image.Point

func drawCell(x, y int, cell int8) {
	var c uint8
	if isAlive(cell) {
		c = 1
	} else {
		c = 0
	}

	x = 3 * (x + origin.X)
	y = 3 * (y + origin.Y)

	display.SetColorIndex(x, y, c)
	display.SetColorIndex(x+1, y, c)
	display.SetColorIndex(x, y+1, c)
	display.SetColorIndex(x+1, y+1, c)
}
