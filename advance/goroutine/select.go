package goroutine

import (
	"fmt"
	"math/rand"
	"time"
)

func createWorker() chan int {
	var result = make(chan int)
	go func(c chan int) {
		for {
			fmt.Printf("received value: %d\n", <-c)
		}
	}(result)
	return result
}

func DoSelect() {
	var c1, c2 = generator(), generator()
	var worker = createWorker()
	var values []int
	for {
		var activeWorker chan int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.Tick(500 * time.Microsecond):
			fmt.Println("queue length:", len(values))
		}
	}
}

func generator() chan int {
	var result = make(chan int)
	go func() {
		var i int
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Microsecond)
			result <- i
			i++
		}
	}()
	return result
}
