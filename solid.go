package main

import "strings"

func solidPattern() string {
	return strings.Repeat(strings.Repeat("O", patternSize)+"\n", patternSize)
}
