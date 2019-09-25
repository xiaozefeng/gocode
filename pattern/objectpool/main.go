package main

import (
	"fmt"
	"github.com/xiaozefeng/gocode/pattern/objectpool/pool"
)

func main() {
	p := pool.New(2)
	for {
		select {
		case obj := <-p:
			obj.Do()
			p <- obj
		default:
			fmt.Println("no more objects")
		}
	}
}
