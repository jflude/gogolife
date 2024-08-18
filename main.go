package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var file = flag.String("file", "life.png", "the path of the output file")
var kind = flag.Int("kind", 0, "the kind of initial pattern")
var gens = flag.Int("gens", 600, "the number of generations to show")
var skip = flag.Int("skip", 0, "the periodicity of generations to display")
var seed = flag.Int64("seed", 0, "the seed for random numbers (0 is randomize)")
var size = flag.Int("size", displayWidth/2, "the initial size of the pattern")
var caption = flag.Bool("caption", false, "show a caption")
var debug = flag.Int("debug", 0, "how much debugging information to output")

var ani animator

func main() {
	flag.Parse()
	*skip++
	if *seed == 0 {
		*seed = time.Now().UnixNano()
	}
	rand.Seed(*seed)
	if *kind == 0 {
		*kind = rand.Intn(3) + 1
	}
	switch ext := strings.ToLower(filepath.Ext(*file)); ext {
	case ".gif":
		ani = new(gifAnimator)
	case ".png":
		ani = new(pngAnimator)
	default:
		quit(1, errors.New("file type must be GIF or PNG"))
	}

	pattern, err := sowPattern(*kind, *size)
	if err != nil {
		quit(2, err)
	}
	if err = importPattern(strings.NewReader(pattern)); err != nil {
		quit(3, err)
	}
	if *debug > 0 {
		fmt.Fprintln(os.Stderr, "kind:", *kind, "size:", *size,
			"procs:", runtime.GOMAXPROCS(0), "seed:", *seed)
	}

	startJob(&normal)
	if err := loop(); err != nil {
		quit(4, err)
	}
}
