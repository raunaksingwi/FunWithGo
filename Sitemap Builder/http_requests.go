package main

import (
	"net/http"

	"golang.org/x/net/html"
)

func getRootNodeOfHTML(url string) (*html.Node, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	rootNode, err := html.Parse(response.Body)
	if err != nil {
		return nil, err
	}
	return rootNode, nil
}
