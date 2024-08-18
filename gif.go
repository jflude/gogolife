package main

import (
	"image"
	"image/gif"
	"io"
)

type gifAnimator struct {
	gif.GIF
}

func (a *gifAnimator) accumulate(img *image.Paletted, delayMillis int) {
	a.Image = append(a.Image, img)
	a.Delay = append(a.Delay, delayMillis/10)
}

func (a *gifAnimator) encode(w io.Writer) error {
	return gif.EncodeAll(w, &a.GIF)
}
