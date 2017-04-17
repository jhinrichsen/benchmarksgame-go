/* The Computer Language Benchmarks Game
   http://benchmarksgame.alioth.debian.org/

   contributed by KP
   modified by fwip
*/

package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
)

type T struct {
	next  *T
	label int
	value int
	mux   sync.Mutex
}

func (w *T) put(res chan int, v int) {
	w.value = v
	if v == 0 {
		res <- w.label
	} else {
		w.mux.Unlock()
	}
}

func (w *T) run(res chan int) {
	for {
		w.mux.Lock()
		w.next.put(res, w.value-1)
		runtime.Gosched()
	}
}

func (w *T) Start(label int, next *T, res chan int) {
	w.label = label
	w.next = next
	w.mux.Lock()
	go w.run(res)
}

const NThreads = 503

func main() {
	n := 1000
	if len(os.Args) > 1 {
		n, _ = strconv.Atoi(os.Args[1])
	}
	runtime.GOMAXPROCS(1)
	fmt.Println(threadRing6(n))
}

func threadRing6(n int) int {
	var res = make(chan int)
	var channels [NThreads]T
	for i := range channels {
		channels[i].Start(i+1, &channels[(i+1)%NThreads], res)
	}
	channels[0].put(res, n)
	return <-res
}
