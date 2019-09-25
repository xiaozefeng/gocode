package memory

import (
	"fmt"
	"io"
)

type Storage struct {
}

func (s Storage) Open(string) (io.ReadWriteCloser, error) {
	fmt.Println("new memory storage")
	return nil, nil
}

func NewStorage() Storage {
	return Storage{}
}
