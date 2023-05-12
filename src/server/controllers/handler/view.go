package handler

import (
	"net/http"
	"text/template"
)

func View(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./frontend/public/index.html"))
	t.Execute(w, nil)
}
