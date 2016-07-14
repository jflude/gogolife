package main

import (
	"runtime"
	"sync"
)

type job interface {
	start()
	work(id int, r *row)
	finish()
}

var stride = runtime.GOMAXPROCS(0)
var wg sync.WaitGroup

func startJob(j job) {
	j.start()
	r := rows.next

	for i := 0; i < stride; i++ {
		wg.Add(1)
		go runJob(j, r, i)
		if r = r.next; r == &rows {
			break
		}
	}

	wg.Wait()
	j.finish()
}

func runJob(j job, r *row, id int) {
	for {
		j.work(id, r)
		for i := 0; i < stride; i++ {
			if r = r.next; r == &rows {
				wg.Done()
				return
			}
		}
	}
}
