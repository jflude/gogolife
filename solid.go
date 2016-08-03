package main

import "strings"

func solidPattern() string {
	line := strings.Repeat("O", patternSize) + "\n"
	return strings.Repeat(line, patternSize)
}
