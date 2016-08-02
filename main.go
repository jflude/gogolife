package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxGenerations = 200

func main() {
	var kind int
	var err error
	if len(os.Args) > 1 {
		if kind, err = strconv.Atoi(os.Args[1]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	pattern := sowPattern(kind)
	if err = importPattern(strings.NewReader(pattern)); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	startJob(&normal)

	for {
		accumulate()
		if generation >= maxGenerations || !generate() {
			break
		}
	}

	if err = encode("life.gif"); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}
