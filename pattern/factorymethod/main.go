package main

import (
	"github.com/xiaozefeng/gocode/pattern/factorymethod/store"
)

func main() {
	s := store.NewStore(store.MemoryStorage)
	_, _ = s.Open("file")
	//f, _ := s.Open("file")
	//defer f.Close()
	//f.Write([]byte("data"))
}
