package main

import (
	"fmt"
	"os"
	"strings"
)

const maxGenerations = 200

func main() {
	pattern := sowPattern()
	if err := importPattern(strings.NewReader(pattern)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	startJob(&normal)

	for {
		accumulate()
		if generation >= maxGenerations || !generate() {
			break
		}
	}

	if err := encode("life.gif"); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
