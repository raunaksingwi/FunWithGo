package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/pkg/profile"
)

var domain string

var seen safeMap

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	seen.links = make(map[string]voidType)
	var maxDepth int
	flag.StringVar(&domain, "rootWebsite", "", "address of root website to start parsing")
	flag.IntVar(&maxDepth, "maxDepth", 3, "Maximum depth to traverse")
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
