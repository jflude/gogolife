package main

import "errors"

const density = 50

func sowPattern(kind int, size int) (string, error) {
	switch kind {
	case 1:
		return randomPattern(size), nil
	case 2:
		return solidPattern(size), nil
	case 3:
		return agarPattern(size), nil
	default:
		return "", errors.New("kind must be in the range 1...3")
	}
}
