package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const maxGenerations = 240

func main() {
	var kind = flag.Int("kind", 0, "the kind of initial pattern")
	var size = flag.Int("size", displayWidth/2, "the initial size of the pattern")
	var seed = flag.Int64("seed", time.Now().UnixNano(), "the seed for random numbers")
	flag.Parse()

	rand.Seed(*seed)
	if *kind == 0 {
		*kind = rand.Intn(3) + 1
	}

	pattern, err := sowPattern(*kind, *size)
	if err != nil {
		quit(1, err)
	}

	if err = importPattern(strings.NewReader(pattern)); err != nil {
		quit(2, err)
	}

	fmt.Println("kind:", *kind, "size:", *size, "seed:", *seed)
	startJob(&normal)

	var ani gifAnimator
	for {
		ani.accumulate()
		if generation >= maxGenerations || !generate() {
			break
		}
	}

	if err = ani.encode("life"); err != nil {
		quit(3, err)
	}
}

func quit(code int, err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(code)
}
