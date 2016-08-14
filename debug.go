package main

import (
	"fmt"
	"os"
	"strings"
)

func printWorld() {
	for r := rows.next; r != &rows; r = r.next {
		var s string
		if r.spans.next == &r.spans {
			s = "()"
		} else {
			for sp := r.spans.next; sp != &r.spans; sp = sp.next {
				off := sp.left
				if sp != r.spans.next {
					off -= sp.prev.right
				}

				if off > 0 {
					s += strings.Repeat(" ", off)
				}

				for _, a := range sp.cells {
					if isAlive(a) {
						s += "O"
					} else {
						s += "."
					}
				}
			}
		}

		fmt.Fprintf(os.Stderr, "%3d (%3d, %3d) -> %s\n",
			r.y, r.spans.next.left, r.spans.next.right, s)
	}

	fmt.Println()
}

func printChanges() {
	for r := rows.next; r != &rows; r = r.next {
		var s, s2 string
		if r.spans.next == &r.spans {
			s = "()"
		} else {
			for sp := r.spans.next; sp != &r.spans; sp = sp.next {
				s += fmt.Sprintf(" -> (%3d,%3d: %v)",
					sp.left, sp.right, sp.cells)
			}
		}

		for _, chg := range r.changes {
			s2 += fmt.Sprintf("%d(%d) ", chg.sp.left, chg.x)
		}

		fmt.Fprintf(os.Stderr, "%3d%-12s <%s>\n", r.y, s, strings.TrimSpace(s2))
	}

	fmt.Fprintln(os.Stderr)
}

func printEffects() {
	for r := rows.next; r != &rows; r = r.next {
		var s, s2 string
		if r.spans.next == &r.spans {
			s = "()"
		} else {
			for sp := r.spans.next; sp != &r.spans; sp = sp.next {
				s += fmt.Sprintf(" -> (%3d,%3d: %v)",
					sp.left, sp.right, sp.cells)
			}
		}

		for _, eff := range r.effects {
			s2 += fmt.Sprintf("%d(%d) ", eff.x, eff.delta)
		}

		fmt.Fprintf(os.Stderr, "%3d%-12s {%s}\n", r.y, s, strings.TrimSpace(s2))
	}

	fmt.Fprintln(os.Stderr)
}
