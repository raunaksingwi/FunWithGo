package main

import (
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func getAllLinks(node *html.Node, nextLevel *safeMap, wg *sync.WaitGroup) {
	if node == nil {
		return
	}

	if node.Type == html.ElementNode && node.Data == "a" {
		href := node.Attr[0].Val
		if !seen.lookUp(href) {
			if strings.HasPrefix(href, domain) || strings.HasPrefix(href, "/") {
				seen.add(href)
				nextLevel.add(href)
			}
		}
	}

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		wg.Add(1)
		go getAllLinks(n, nextLevel, wg)
	}
	wg.Done()
}

func traverseLinks(startPage string, maxDepth int) {
	var wg sync.WaitGroup
	var nextLevel safeMap
	nextLevel.links = make(map[string]voidType)

	var toTraverse safeMap
	toTraverse.links = make(map[string]voidType)
	nextLevel.add(startPage)

	for maxDepth > 0 && len(nextLevel.links) > 0 {
		toTraverse.links = nextLevel.links
		nextLevel.links = make(map[string]voidType)

		for url := range toTraverse.links {
			if !strings.HasPrefix(url, domain) {
				url = domain + url
			}
			node, err := getRootNodeOfHTML(url)
			if err != nil {
				panic(err)
			}
			wg.Add(1)
			go getAllLinks(node, &nextLevel, &wg)
		}
		wg.Wait()
		maxDepth--
	}

}
