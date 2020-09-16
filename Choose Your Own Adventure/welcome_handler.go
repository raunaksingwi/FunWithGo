package main

import (
	"html/template"
	"log"
	"net/http"
)

type welcomeHandler struct {
}

func (welcome welcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(htmlTemplate)
	if err != nil {
		log.Println("Failed to parse html template; err=", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, storyMap["intro"])
}
