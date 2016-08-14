package main

import (
	"fmt"
	"os"
)

func quit(code int, err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	os.Exit(code)
}
