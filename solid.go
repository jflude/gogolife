package main

import "strings"

func solidPattern(size int) string {
	return "0 0\n" + strings.Repeat(strings.Repeat("O", size)+"\n", size)
}
