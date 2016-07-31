package main

import (
	"image"
	"image/color"
)

var origin = image.Pt(150, 150)
var display = image.NewPaletted(image.Rect(0, 0, 999, 999),
	color.Palette{color.White, color.Black})

func drawCell(x, y int, cell int8) {
	var c uint8
	if isAlive(cell) {
		c = 1
	} else {
		c = 0
	}

	x = 3*x + origin.X
	y = 3*y + origin.Y

	i := (y-display.Rect.Min.Y)*display.Stride + x - display.Rect.Min.X

	display.Pix[i] = c
	display.Pix[i+1] = c

	i += display.Stride

	display.Pix[i] = c
	display.Pix[i+1] = c
}
