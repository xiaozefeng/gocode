package store

import (
	"github.com/xiaozefeng/gocode/pattern/factorymethod/store/disk"
	"github.com/xiaozefeng/gocode/pattern/factorymethod/store/memory"
	"github.com/xiaozefeng/gocode/pattern/factorymethod/store/temp"
	"io"
)

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

func NewStore(storageType StorageType) Store {
	switch storageType {
	case MemoryStorage:
		return memory.NewStorage()
	case DiskStorage:
		return disk.NewStorage()
	default:
		return temp.NewStorage()
	}
}
