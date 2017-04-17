package main

const threads = 503

func main() {
}

func mkChannels() (channels [threads]chan int) {
	for i := 0; i < threads; i++ {
		channels[i] = make(chan int)
	}
	return
}

func threadRing(n int) int {
	rc := make(chan int)
	cs := mkChannels()
	for i := 0; i < n; i++ {
		go pass(n, cs[i%threads], cs[(i+1)%threads], rc)
	}
	// log.Printf("Write %v to channel %v\n", n, 0)
	cs[0] <- 0
	return 1 + <-rc
}

func pass(n int, prev <-chan int, succ chan<- int, rc chan int) {
	i := <-prev
	// log.Printf("read %v\n", i)
	i++
	if i == n {
		// log.Printf("write to result channel: %v\n", i)
		rc <- n % threads
		// log.Println("closing result channel")
		close(rc)
	} else {
		// log.Printf("write %v\n", i)
		succ <- i
	}
}
