package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func getRootNodeOfHTML(url string) (*html.Node, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)
	r := strings.NewReader(pageContent)

	rootNode, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return rootNode, nil
}
