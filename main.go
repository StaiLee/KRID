package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Valeur string
}

const port = 1709

var Templateshome = template.Must(template.ParseFiles("./static/templates/index.html"))

func main() {
	http.HandleFunc("./index", homeHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println("http://localhost:1709 - server started on port", port)
	http.ListenAndServe(":1709", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Valeur: "Home page"}
	err := Templateshome.ExecuteTemplate(w, "index.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
