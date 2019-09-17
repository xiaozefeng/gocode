package errorhanding

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println(err)
		} else {
			panic(r)
		}
	}()

	//panic(errors.New("this is a error"))

	var a = 0
	var b = 1 / a
	fmt.Println(b)
}
