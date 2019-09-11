package slices

import "fmt"

func createSlice() {
	arr := [6]int{0, 1, 2, 3, 4, 5}

	fmt.Println("arr[2:5]", arr[2:5])
	fmt.Println("arr[:5]", arr[:5])
	fmt.Println("arr[2:]", arr[2:])
	fmt.Println("arr[:]", arr[:])

	// reSlice
	s1 := arr[2:]
	printSlice("s1", s1)
	s2 := s1[2:3]
	printSlice("s2", s2)

	// extend slice
	s3 := s2[:2]
	printSlice("s3", s3)

	// append
	s4 := append(s3, 10)
	s5 := append(s4, 11)
	printSlice("s4", s4)
	printSlice("s5", s5)
}

func printSlice(name string, s []int) {
	fmt.Printf("%s:%v, len:%d, cap:%d\n", name, s, len(s), cap(s))
}
