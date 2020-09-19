package main

import "sync"

type voidType struct{}

var void voidType

type safeMap struct {
	links map[string]voidType
	mux   sync.Mutex
}

func (s *safeMap) add(link string) {
	s.mux.Lock()
	s.links[link] = void
	s.mux.Unlock()
}

func (s *safeMap) lookUp(link string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	_, ok := s.links[link]
	return ok
}
