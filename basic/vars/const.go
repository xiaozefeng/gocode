package vars

import "fmt"

const (
	FILENAME = "abc.txt"
	PI       = 3.1415
)

func printConst() {
	fmt.Println(FILENAME)
	fmt.Println(PI)
}

func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)
}

func fileTypeSize() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b, kb, mb, gb, tb, pb)
}
