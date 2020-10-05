package singleton

import "sync"

type singletonLock struct {
	count int
	sync.RWMutex
}

var instanceLock singletonLock

func GetInstanceLock() *singletonLock {
	return &instanceLock
}

func (s *singletonLock) AddOneLock() {
	s.Lock()
	defer s.Unlock()
	s.count++
}

func (s *singletonLock) GetCountLock() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}
