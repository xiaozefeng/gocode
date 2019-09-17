package goroutine

import (
	"fmt"
)

func channelDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go work(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
}

func work(id int, c chan int) {
	for {
		fmt.Printf("Worker:%d received %c\n", id, <-c)
	}
}

