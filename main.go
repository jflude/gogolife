package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const maxGenerations = 240

func main() {
	seed := time.Now().UnixNano()
	fmt.Println("Seed is", seed)
	rand.Seed(seed)

	var kind int
	var err error

	switch len(os.Args) {
	case 3:
		if patternSize, err = strconv.Atoi(os.Args[2]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fallthrough
	case 2:
		if kind, err = strconv.Atoi(os.Args[1]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case 1:
		kind = rand.Intn(3)
	}

	pattern := sowPattern(kind)
	if err = importPattern(strings.NewReader(pattern)); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	startJob(&normal)

	var ani gifAnimator
	for {
		ani.accumulate()
		if generation >= maxGenerations || !generate() {
			break
		}
	}

	if err = ani.encode("life.gif"); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}
