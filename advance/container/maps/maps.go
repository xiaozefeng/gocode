package maps

import "fmt"

var (
	m1 map[string]string
	m2 = map[string]string{
		"name": "jack",
		"age":  " 18",
	}
)

func createMap() {
	fmt.Println("m1", m1)

	fmt.Println("m2", m2)

	var m3 = make(map[string]string)
	m3["name"] = "jack"
	m3["hobbies"] = "吃饭，睡觉，打代码"
	fmt.Println("m3", m3)
}

func rangeMap() {
	for k,v :=range m2 {
		fmt.Printf("k=%s,v=%s\n", k, v)
	}
}

