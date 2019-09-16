package functional

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", f())
	}
}

func TestIntGen(t *testing.T) {
	f := Fibonacci()
	Read(f)
}


