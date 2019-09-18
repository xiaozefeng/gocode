package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAtomicInt(t *testing.T) {
	var ai = &AtomicInt{
		Value:10,
		Lock: sync.Mutex{},
	}

	ai.Increment()
	go func() {
		ai.Increment()
	}()

	time.Sleep(5 * time.Microsecond)
	fmt.Println("value:",ai.Get())
}

