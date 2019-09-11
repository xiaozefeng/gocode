package array

import "fmt"

func createArray() {
	var arr1 [5] int
	var arr2 = [5]int{2, 3, 4, 5, 6}
	var arr3 = [...] int{2, 4, 6, 8, 10}
	var grid [4][5] int
	fmt.Println("arr1", arr1)
	fmt.Println("arr2", arr2)
	fmt.Println("arr3", arr3)
	fmt.Println("grid", grid)
}

func rangeArray() {
	var arr = [5]int{2, 3, 4, 5, 6}
	fmt.Println("range way 1")
	for i := range arr{
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	fmt.Println("range way 2")
	for i, v :=range arr{
		fmt.Printf("index:%d, value:%d\n", i, v)
	}
	fmt.Println()

	fmt.Println("range way 3")
	for _, v:= range arr{
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
