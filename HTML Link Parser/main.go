package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type link struct {
	Href, Text string
}

var allLinks []link

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

func getAllLinks(node *html.Node) {
	if node == nil {
		return
	}

	if node.Type == html.ElementNode && node.Data == "a" {
		newLink := link{
			Href: node.Attr[0].Val,
		}
		populateNodeText(node, &newLink)
		allLinks = append(allLinks, newLink)

	}
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		getAllLinks(n)
	}
}

func populateNodeText(node *html.Node, l *link) {
	if node.Type == html.TextNode {
		l.Text += strings.TrimSpace(node.Data)
	}
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		populateNodeText(n, l)
	}

}

func main() {
	url := "https://google.org"
	rootNode, err := getRootNodeOfHTML(url)
	if err != nil {
		log.Fatalf("Failed to parse the URL; url: %s; err: %v ", url, err)
	}

	getAllLinks(rootNode)

	fmt.Println("All links")
	fmt.Println("=====================")
	for _, mylink := range allLinks {
		fmt.Printf("%s -> %s \n", mylink.Href, mylink.Text)
	}

}
