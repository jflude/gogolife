package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const maxAccumulations = 240

var debug = flag.Int("debug", 0, "how much debugging information to output")
var kind = flag.Int("kind", 0, "the kind of initial pattern")
var period = flag.Int("period", 1, "the periodicity of generations to display")
var seed = flag.Int64("seed", time.Now().UnixNano(), "the seed for random numbers")
var size = flag.Int("size", displayWidth/2, "the initial size of the pattern")

func main() {
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

	if *debug > 0 {
		fmt.Fprintln(os.Stderr, "kind:", *kind, "size:", *size, "seed:", *seed)
	}

	startJob(&normal)

	if err := loop(*period * maxAccumulations); err != nil {
		quit(3, err)
	}
}
