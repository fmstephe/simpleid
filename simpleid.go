package simpleid

import (
	"sync"
	"strconv"
)

type IdMaker interface {
	NewId() string
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
	return strconv.FormatInt(i.id,16)
}
