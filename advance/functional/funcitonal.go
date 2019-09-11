package functional

import (
	"fmt"
	"reflect"
	"runtime"
)

// function as a parameter
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling funciton %s with args"+
		"(%d, %d)", opName, a, b)
	return op(a, b)
}

func add(a, b int) int {
	return a + b
}

// 可变参数
func sum(nums ...int) int {
	var sum int
	for _, i := range nums {
		sum += i
	}
	return sum
}

// swap
func swap(a, b int) (int, int) {
	return b, a
}
