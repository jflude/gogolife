package main

type change struct {
	sp *span
	x  int
}

type effect struct {
	x     int
	delta int8
}

type ChangeSlice []change

func (cs ChangeSlice) Len() int {
	return len(cs)
}

func (cs ChangeSlice) Less(i, j int) bool {
	return cs[i].x < cs[j].x
}

func (cs ChangeSlice) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}
