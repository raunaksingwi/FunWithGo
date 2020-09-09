package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var jsonFile = "./resources/gopher.json"
var httpLayout = "./static/layout.html"
var storyMap map[string]Story

func main() {
	var err error
	storyMap, err = getStoriesFromJSON(jsonFile)
	if err != nil {
		log.Fatalln("Failed to load json; err=", err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/{story}", displayStory)
	fmt.Println("Starting server on localhost:8080")
	log.Fatalln(http.ListenAndServe(":8080", router))

}

func displayStory(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(httpLayout)
	if err != nil {
		log.Println("Failed to parse html template; err=", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	tmpl.Execute(w, storyMap[vars["story"]])
}
