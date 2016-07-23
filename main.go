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

	startJob(&normalize)

	for generate(1000) {
		// empty
	}
}
