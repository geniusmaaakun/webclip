package handler

import (
	"html/template"
	"log"
	"net/http"
)

func View(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./frontend/build/index.html"))
	err := t.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
