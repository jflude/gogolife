package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	pattern := sowPattern()
	if err := importPattern(strings.NewReader(pattern)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	startJob(&normal)

	for generate(10) {
		if err := encode(fmt.Sprintf("img%06d", generation)); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	}
}
