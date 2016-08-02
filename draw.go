package main

import (
	"image"
	"image/color"
)

const (
	displayWidth  = 333
	displayHeight = 333
)

var display = image.NewPaletted(image.Rect(0, 0, 3*displayWidth, 3*displayHeight),
	color.Palette{color.White, color.Black})

var origin image.Point

func drawCell(x, y int, cell int8) {
	var c uint8
	if isAlive(cell) {
		c = 1
	} else {
		c = 0
	}

	x = 3*x + origin.X
	y = 3*y + origin.Y

	display.SetColorIndex(x, y, c)
	display.SetColorIndex(x+1, y, c)
	display.SetColorIndex(x, y+1, c)
	display.SetColorIndex(x+1, y+1, c)
}
