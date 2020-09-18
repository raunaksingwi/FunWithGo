package main

import (
	"flag"
	"fmt"
	"sort"
	"sync"
)

var domain string

type voidType struct{}

var void voidType

type seenStruct struct {
	links map[string]voidType
	mux   sync.Mutex
}

func (s *seenStruct) add(link string) {
	s.mux.Lock()
	s.links[link] = void
	s.mux.Unlock()
}

func (s *seenStruct) lookUp(link string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	_, ok := s.links[link]
	return ok
}

var seen seenStruct

func main() {
	seen.links = make(map[string]voidType)
	var maxDepth int
	flag.StringVar(&domain, "rootWebsite", "", "address of root website to start parsing")
	flag.IntVar(&maxDepth, "maxDept", 3, "Maximum depth to traverse")
	flag.Parse()
	seen.add(domain)
	traverseLinks(domain, maxDepth)

	keys := make([]string, len(seen.links))
	i := 0
	for k := range seen.links {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key)
	}

}
