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
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("./static/assets/css/"))))
	http.Handle("/assets/js/", http.StripPrefix("/assets/js/", http.FileServer(http.Dir("./static/assets/js/"))))
	http.Handle("/assets/images/", http.StripPrefix("/assets/images/", http.FileServer(http.Dir("./static/assets/images/"))))
	http.HandleFunc("/index.html", homeHandler)
	fmt.Printf("Started server successfully on http://localhost:1709/home\n")
	http.ListenAndServe(":1709", nil)
}
