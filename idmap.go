package simpleid

import (
	"errors"
	"fmt"
	"sync"
)

type IdMap interface {
	Contains(string) bool
	Add(string,interface{}) error
	Remove(string)
	Get(string) interface{}
}

type lockedIdMap struct {
	ids  map[string] interface{}
	lock sync.Mutex
}

func NewIdMap() IdMap {
	s := new(lockedIdMap)
	s.ids = make(map[string]interface{})
	return s
}

func (s *lockedIdMap) Contains(id string) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, present := s.ids[id]
	return present
}

func (s *lockedIdMap) Add(id string, val interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, present := s.ids[id]; present {
		return errors.New(fmt.Sprintf("id %s already present in id map", id))
	}
	s.ids[id] = val
	return nil
}

func (s *lockedIdMap) Remove(id string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.ids, id)
}

func (s *lockedIdMap) Get(id string) interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.ids[id]
}
