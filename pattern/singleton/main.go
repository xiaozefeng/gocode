package main

import (
	"fmt"
	"sync"
)

type single map[string]string

var (
	once     sync.Once
	instance single
)

func New() single {
	once.Do(func() {
		instance = make(single)
	})
	return instance
}

func main() {
	s := New()
	s["this"] = "that"

	s1 := New()
	fmt.Println("this is", s1["this"])
}
