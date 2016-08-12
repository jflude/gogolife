package main

type animator interface {
	accumulate()
	encode(filename string) error
}
