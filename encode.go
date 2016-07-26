package main

import (
	"image/jpeg"
	"os"
)

func encode(name string) error {
	f, err := os.Create(name + ".jpg")
	if err != nil {
		return err
	}

	if err := jpeg.Encode(f, display, &jpeg.Options{Quality: 100}); err != nil {
		f.Close()
		return err
	}

	return f.Close()
}
