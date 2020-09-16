package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var htmlTemplate = "./static/layout.html"

var storyMap map[string]Story

func main() {

	var jsonFile = "./resources/gopher.json"

	var err error
	storyMap, err = getStoriesFromJSON(jsonFile)
	if err != nil {
		log.Fatalln("Failed to load json; err=", err)
	}

	router := mux.NewRouter()

	router.Handle("/", welcomeHandler{})
	router.Handle("/{story}", storyHandler{})

	log.Println("Starting server on localhost:8080")
	log.Fatalln(http.ListenAndServe(":8080", router))
}
