package simpleid

import (
	"sync"
)

type IdMaker struct {
	id   int64
	lock sync.Mutex
}

func New() *IdMaker {
	return new(IdMaker)
}

func (i *IdMaker) NewId() int64 {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.id++
	return i.id
}
