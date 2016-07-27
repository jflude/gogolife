package main

import (
	"image"
	"image/color"
	"image/draw"
)

var display = image.NewRGBA(image.Rect(0, 0, 1000, 1000))
var offset = image.Pt(400, 400)

func init() {
	draw.Draw(display, display.Bounds(), &image.Uniform{C: color.White}, image.ZP, draw.Src)
}

func drawCell(x, y int, cell int8) {
	var c color.Color
	if isAlive(cell) {
		c = color.Black
	} else {
		c = color.White
	}

	x = 3*x + offset.X
	y = 3*y + offset.Y

	display.Set(x, y, c)
	display.Set(x+1, y, c)
	display.Set(x, y+1, c)
	display.Set(x+1, y+1, c)
}
