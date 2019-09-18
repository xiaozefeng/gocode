package goroutine

import "sync"

type AtomicInt struct {
	Value int
	Lock sync.Mutex
}

func (a *AtomicInt) Increment(){
	a.Lock.Lock()
	defer a.Lock.Unlock()
	a.Value ++
}

func (a * AtomicInt) Get() int {
	a.Lock.Lock()
	defer a.Lock.Unlock()
	return a.Value
}

