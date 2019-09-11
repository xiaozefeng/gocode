package functional

import (
	"fmt"
	"testing"
)

func TestApply(t *testing.T) {
	r := apply(add, 1, 2)
	fmt.Println(r)
}

func TestSum(t *testing.T) {
	r := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(r)
}

func TestSwap(t *testing.T) {
	var a, b = 3, 4
	fmt.Println("a,b", a, b)
	a, b = swap(a, b)
	fmt.Println("after swap , a,b ", a, b)
}
