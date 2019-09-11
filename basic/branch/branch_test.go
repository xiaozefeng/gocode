package branch

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	readFile()
	fmt.Println("")
	readFile2()
}

func TestEval(t *testing.T) {
	fmt.Println(eval(1,2, "+"))
	fmt.Println(eval(1,2, "-"))
	fmt.Println(eval(1,2, "*"))
	fmt.Println(eval(1,2, "/"))
}

func TestGrade(t *testing.T) {
	fmt.Println(grade(100))
	fmt.Println(grade(86))
	fmt.Println(grade(52))
	fmt.Println(grade(67))
	fmt.Println(grade(77))
	fmt.Println(grade(33))
}

