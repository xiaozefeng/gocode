package functional

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Fibonacci() IntGen{
	var a, b = 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 为函数实现接口
type IntGen func() int

func (g IntGen) Read(p []byte) (n int, err error) {
	var next = g()
	if next > 10000 {
		return  0,io.EOF
	}
	var s = fmt.Sprintf("%d ",next)
	return strings.NewReader(s).Read(p)
}

func Read(r io.Reader){
	scanner := bufio.NewScanner(r)
	for scanner.Scan(){
		fmt.Printf("%v", scanner.Text())
	}
}


