package types

import (
	"fmt"
	"math"
)

func printType() {
	var a = true
	var b = "string"
	var c int
	c = 123
	var d float64
	d = 1.345
	var e byte
	var f rune
	fmt.Println(a, b, c, d)
	fmt.Println(e, f)
}

func typeConvert() {
	var a float64 = 1.2
	var b int = int(a)
	// 精度损失
	fmt.Println("convert float64 to int:", b)

	a = float64(b)
	fmt.Println("convert int to float64:", a)

	// string to byte
	var c string = "hello"
	var bytes []byte = []byte(c)
	fmt.Println("string to bytes:", bytes)

	// byte to string
	c = string(bytes)
	fmt.Println("bytes to string:", c)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println("c=", c)
}
