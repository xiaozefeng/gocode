package main

import "log"

type Object func(int) int

func LogDecorator(fn Object) Object {
	return func(i int) int {
		log.Println("Starting the execution with the integer ", i)
		result := fn(i)
		log.Println("Execution is completed with the result", result)
		return result
	}
}

func main() {
	f := LogDecorator(Double)
	f(5)
}

func Double(n int) int {
	return n * 2
}
