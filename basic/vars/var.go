package vars

import "fmt"

var (
	aaa = "aaa"
	bbb = 222
	ccc = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "hello"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, d = 3, 4, true, 1.34
	var s = "hello"
	fmt.Println(a, b, c, d, s)
}

func variableShorter() {
	a := 1
	b := true
	c := "hello"
	fmt.Println(a, b, c)
}

func packageVariable() {
	fmt.Println(aaa, bbb, ccc)
}
