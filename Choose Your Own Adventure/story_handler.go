package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type storyHandler struct {
}

func (s storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(htmlTemplate)
	if err != nil {
		log.Println("Failed to parse html template; err=", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	if _, ok := storyMap[vars["story"]]; ok {
		tmpl.Execute(w, storyMap[vars["story"]])
	}

}
