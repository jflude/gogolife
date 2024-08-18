package main

import (
	"image"
	"io"

	"github.com/kettek/apng"
)

type pngAnimator struct {
	apng.APNG
}

func (a *pngAnimator) accumulate(img *image.Paletted, delayMillis int) {
	a.Frames = append(a.Frames, apng.Frame{
		Image:            img,
		DelayNumerator:   uint16(delayMillis),
		DelayDenominator: 1000,
	})
}

func (a *pngAnimator) encode(w io.Writer) error {
	return apng.Encode(w, a.APNG)
}
