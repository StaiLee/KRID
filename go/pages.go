package module

//import (
//	"net/http"
//	"text/template"
//)

//var TemplateHome = template.Must(template.ParseFiles("./static/templates/index.html"))

//func homeHandler(w http.ResponseWriter, r *http.Request) {
//	p := Page{Valeur: "Home page"}
//	err := TemplateHome.ExecuteTemplate(w, "index.html", p)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
