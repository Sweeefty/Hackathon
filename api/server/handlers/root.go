package handlers

import (
	"net/http"
	"text/template"
)

func Root(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/html/index.html")) //We link the template and the html file
	tmpl.Execute(w, nil)
}
