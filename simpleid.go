package simpleid

import (
	"strconv"
	"sync"
)

type IdMaker interface {
	NewId() string
	Id() int64
}

type lockedMaker struct {
	id   int64
	lock sync.Mutex
}

func NewIdMaker() IdMaker {
	return new(lockedMaker)
}

func (i *lockedMaker) NewId() string {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.id++
	return strconv.FormatInt(i.id, 16)
}

func (i *lockedMaker) Id() int64 {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.id++
	return i.id
}
