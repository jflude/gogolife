package main

import "image"

type attrib struct {
	left, right, pop int
	removable        []*row
}

type tidyJob struct {
	attribs []attrib
}

var tidy tidyJob

func (j *tidyJob) start() {
	j.attribs = make([]attrib, stride)

	for i := 0; i < stride; i++ {
		j.attribs[i].left = maxInt
		j.attribs[i].right = minInt
	}
}

func (j *tidyJob) work(id int, r *row) {
	if pruneSpans(r) {
		j.attribs[id].removable = append(j.attribs[id].removable, r)
		return
	}

	divideSpans(r)
	clipSpans(r)
	combineSpans(r)

	r.effects = r.effects[:0]

	if r.spans.next.left < j.attribs[id].left {
		j.attribs[id].left = r.spans.next.left
	}

	if r.spans.prev.right > j.attribs[id].right {
		j.attribs[id].right = r.spans.prev.right
	}

	j.attribs[id].pop += r.pop
}

func (j *tidyJob) finish() {
	if rows.next == &rows {
		extent = image.Rectangle{}
		return
	}

	extent = image.Rectangle{
		Min: image.Point{maxInt, rows.next.y},
		Max: image.Point{minInt, rows.prev.y},
	}

	for i := 0; i < stride; i++ {
		if j.attribs[i].left < extent.Min.X {
			extent.Min.X = j.attribs[i].left
		}

		if j.attribs[i].right > extent.Max.X {
			extent.Max.X = j.attribs[i].right
		}

		population += j.attribs[i].pop

		for _, r := range j.attribs[i].removable {
			removeRow(r)
		}
	}
}
