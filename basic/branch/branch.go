package branch

import (
	"fmt"
	"io/ioutil"
)

const filename = "./branch.go"

func readFile() {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%s\n", bytes)
	}
}

func readFile2() {
	if bytes, err := ioutil.ReadFile(filename); err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%s\n", bytes)
	}
}

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation:" + op)
	}
}

func grade(score int) string {
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score:%d", score))
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 60:
		return "C"
	default:
		return "D"
	}
}
