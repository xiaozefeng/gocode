package functional

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("1+...+%d = %d\n", i, a(i))
	}
}

func TestAdder2(t *testing.T) {
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0+1+...+%d =%d\n", i, s)
	}
}
