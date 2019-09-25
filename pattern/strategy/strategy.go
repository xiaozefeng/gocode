package main

import "fmt"

type Operator interface {
	Apply(leftValue int, rightValue int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue int, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

type Addition struct {
}

func (Addition) Apply(leftValue int, rightValue int) int {
	return leftValue + rightValue
}

func main() {
	op := Operation{Addition{}}
	r := op.Operate(3, 5)
	fmt.Println(r)

	op.Operator = Multiplication{}
	r = op.Operate(4, 5)
	fmt.Println(r)

}

type Multiplication struct {
}

func (Multiplication) Apply(leftValue int, rightValue int) int {
	return leftValue * rightValue
}
