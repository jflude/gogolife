package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
)

const maxAccumulations = 240

var debug = flag.Int("debug", 0, "how much debugging information to output")
var file = flag.String("file", "life.gif", "the name of the output file")
var kind = flag.Int("kind", 0, "the kind of initial pattern")
var skip = flag.Int("skip", 0, "the periodicity of generations to display")
var seed = flag.Int64("seed", 0, "the seed for random numbers (0 is randomize)")
var size = flag.Int("size", displayWidth/2, "the initial size of the pattern")

func main() {
	flag.Parse()
	if *seed == 0 {
		*seed = time.Now().UnixNano()
	}
	rand.Seed(*seed)

	*skip++
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
		fmt.Fprintln(os.Stderr, "kind:", *kind, "size:", *size,
			"procs:", runtime.GOMAXPROCS(0), "seed:", *seed,)
	}

	startJob(&normal)

	var ani gifAnimator
	if err := loop(&ani, *file, *skip*maxAccumulations); err != nil {
		quit(3, err)
	}
}
