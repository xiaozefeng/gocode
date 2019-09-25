package temp

import (
	"fmt"
	"io"
)

type Storage struct {
}

func (s Storage) Open(string) (io.ReadWriteCloser, error) {
	fmt.Println("new temp storage")
	return nil, nil
}

func NewStorage() Storage {
	return Storage{}
}
