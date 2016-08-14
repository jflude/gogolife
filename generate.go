package main

var generation, population int

func generate() bool {
	population = 0

	startJob(&apply)
	startJob(&calc)
	startJob(&tidy)

	generation++
	return rows.next != &rows
}
