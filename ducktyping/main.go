package main

import (
	"fmt"
	"github.com/xiaozefeng/gocode/ducktyping/mock"
	"github.com/xiaozefeng/gocode/ducktyping/realretriever"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{Content: "This is fake retriever"}
	inspect(r)
	r = &realretriever.Retriever{UserAgent: "Mozilla/5.0"}
	inspect(r)

	// type assertion
	realRetriever := r.(*realretriever.Retriever)
	fmt.Println("UserAgent:", realRetriever.UserAgent)

	// type assertion 2
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println("Contents:", mockRetriever.Content)
	} else {
		fmt.Println("not a mock retriever")
	}
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Printf("Content: %s\n", v.Content)
	case *realretriever.Retriever:
		fmt.Printf("UserAgent:%s\n", v.UserAgent)
	}
}
