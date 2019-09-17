package goroutine

import (
	"fmt"
	"sync"
)

type Worker struct {
	In   chan int
	Done func()
}

func done() {
	var channels [10]Worker
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		channels[i] = CreateWorker(i, &wg)
	}

	for i, w := range channels {
		w.In <- 'a' + i
	}

	for i, w := range channels {
		w.In <- 'A' + i
	}

	wg.Wait()
}

func CreateWorker(n int, wg *sync.WaitGroup) Worker {
	var w = Worker{
		In: make(chan int),
		Done: func() {
			wg.Done()
		},
	}
	go func(w Worker, n int) {
		for {
			fmt.Printf("Worker:%d received: %c\n", n, <-w.In)
			w.Done()
		}
	}(w, n)
	return w
}
