package main

import (
	"fmt"
	"github.com/xiaozefeng/gocode/advance/dorpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)

	var result float64
	if err = client.Call("Service.Div", dorpc.ServiceArgs{A: 10, B: 3}, &result); err != nil {
		panic(err)
	} else {
		fmt.Println("result:", result)
	}
	if err = client.Call("Service.Div", dorpc.ServiceArgs{A: 10, B: 0}, &result); err != nil {
		panic(err)
	} else {
		fmt.Println("result:", result)
	}
}
