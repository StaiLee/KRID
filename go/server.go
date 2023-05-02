package module

import (
	"fmt"
	"net/http"
	"text/template"
)

var TemplateHome = template.Must(template.ParseFiles("./static/templates/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Valeur: "Home page"}
	err := TemplateHome.ExecuteTemplate(w, "index.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Start() {
	http.Handle("/static/assets/style/", http.StripPrefix("/static/assets/style/", http.FileServer(http.Dir("./static/assets/style/"))))
	http.Handle("/static/assets/js/", http.StripPrefix("/static/assets/js/", http.FileServer(http.Dir("./static/assets/js/"))))
	http.Handle("/static/assets/images/", http.StripPrefix("/static/assets/images/", http.FileServer(http.Dir("./static/assets/images/"))))
	http.Handle("/static/templates/", http.StripPrefix("/static/templates/", http.FileServer(http.Dir("./static/templates"))))
	http.HandleFunc("/", homeHandler)
	fmt.Printf("Started server successfully on http://localhost:1709\n")
	http.ListenAndServe(":1709", nil)
}
