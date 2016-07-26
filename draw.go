package main

import (
	"image"
	"image/color"
)

var display = image.NewRGBA(image.Rect(0, 0, 1000, 1000))
var offset = image.Pt(400, 400)

func drawCell(x, y int, cell int8) {
	var c color.Color
	if isAlive(cell) {
		c = color.White
	} else {
		c = color.Black
	}

	display.Set(x+offset.X, y+offset.Y, c)
}
