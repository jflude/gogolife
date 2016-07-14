package main

const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

type row struct {
	y       int
	pop     int
	spans   span
	changes []change
	effects []effect
	next    *row
	prev    *row
}

var rows row

func init() {
	rows = *newRow(minInt)
	rows.spans.next = nil
	rows.spans.prev = nil
	rows.next = &rows
	rows.prev = &rows
}

func newRow(y int) *row {
	r := &row{y: y, spans: span{left: maxInt, right: minInt}}
	r.spans.next = &r.spans
	r.spans.prev = &r.spans
	return r
}

func insertBeforeRow(r, next *row) *row {
	r.prev = next.prev
	r.next = next
	next.prev = r
	r.prev.next = r
	return r
}

func insertAfterRow(r, prev *row) *row {
	r.next = prev.next
	r.prev = prev
	prev.next = r
	r.next.prev = r
	return r
}

func removeRow(r *row) {
	r.prev.next = r.next
	r.next.prev = r.prev
}
