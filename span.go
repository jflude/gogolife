package main

const cellChunk = 4

type span struct {
	left  int
	right int
	near  int
	cells []int8
	next  *span
	prev  *span
}

func newSpan(x int) *span {
	x = x &^ (cellChunk - 1)
	return &span{left: x, right: x + cellChunk, cells: make([]int8, cellChunk)}
}

func insertBeforeSpan(sp, next *span) *span {
	sp.prev = next.prev
	sp.next = next
	next.prev = sp
	sp.prev.next = sp
	return sp
}

func insertAfterSpan(sp, prev *span) *span {
	sp.next = prev.next
	sp.prev = prev
	prev.next = sp
	sp.next.prev = sp
	return sp
}

func removeSpan(sp *span) {
	sp.prev.next = sp.next
	sp.next.prev = sp.prev
}
