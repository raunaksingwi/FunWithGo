package main

import (
	"flag"
	"fmt"
	"sort"
	"sync"

	"github.com/pkg/profile"
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
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	seen.links = make(map[string]voidType)
	var maxDepth int
	flag.StringVar(&domain, "rootWebsite", "https://www.calhoun.io", "address of root website to start parsing")
	flag.IntVar(&maxDepth, "maxDepth", 2, "Maximum depth to traverse")
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
