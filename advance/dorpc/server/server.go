package main

import (
	"github.com/xiaozefeng/gocode/advance/dorpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	err := rpc.Register(dorpc.Service{})
	if err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
