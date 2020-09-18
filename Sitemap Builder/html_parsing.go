package main

import (
	"strings"

	"golang.org/x/net/html"
)

func getAllLinks(node *html.Node) []string {
	if node == nil {
		return nil
	}
	var allLinks []string
	if node.Type == html.ElementNode && node.Data == "a" {
		href := node.Attr[0].Val
		if !seen.lookUp(href) {
			if strings.HasPrefix(href, domain) || strings.HasPrefix(href, "/") {
				seen.add(href)
				allLinks = append(allLinks, href)
			}
		}
	}
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		allLinks = append(allLinks, getAllLinks(n)...)
	}
	return allLinks
}

func traverseLinks(startPage string, maxDepth int) {
	var nextLevel []string
	nextLevel = append(nextLevel, startPage)
	for maxDepth > 0 && len(nextLevel) > 0 {
		toTraverse := nextLevel
		nextLevel = nil
		for _, url := range toTraverse {
			if !strings.HasPrefix(url, domain) {
				url = domain + url
			}
			node, err := getRootNodeOfHTML(url)
			if err != nil {
				panic(err)
			}
			links := getAllLinks(node)
			nextLevel = append(nextLevel, links...)
		}
		maxDepth--
	}

}
