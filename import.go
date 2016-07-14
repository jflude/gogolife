package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func importPattern(rdr io.Reader) error {
	scan := bufio.NewScanner(rdr)
	var x, y, count int

	for scan.Scan() {
		count++
		line := strings.TrimSpace(scan.Text())
		if len(line) == 0 {
			continue
		}

		if line[0] == '.' || line[0] == 'O' || line[0] == 'o' {
			var alive []bool
			for _, c := range line {
				alive = append(alive, c != '.')
			}

			placeCells(x, y, alive)
			y++
		} else {
			if n, err := fmt.Sscanf(line, "%d %d\n", &x, &y); err != nil {
				return err
			} else if n != 2 {
				return fmt.Errorf("syntax error [%d]: \"%s\"", count, line)
			}
		}
	}

	return scan.Err()
}
