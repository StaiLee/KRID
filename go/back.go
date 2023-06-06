package module

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// on handle l'integralité des dossier, templates et des differents composants pour notre serveur
func Setup() {
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("./static/assets/css/"))))
	http.Handle("/assets/js/", http.StripPrefix("/assets/js/", http.FileServer(http.Dir("./static/assets/js/"))))
	http.Handle("/assets/images/", http.StripPrefix("/assets/images/", http.FileServer(http.Dir("./static/assets/images/"))))
	http.Handle("/assets/img/", http.StripPrefix("/assets/img/", http.FileServer(http.Dir("./static/assets/img/"))))
	http.Handle("/assets/temp-images/", http.StripPrefix("/assets/temp-images/", http.FileServer(http.Dir("./static/assets/temp-images/"))))
	http.HandleFunc("/templates/login", Login)
	http.HandleFunc("/templates/about", Aboutus)
	http.HandleFunc("/templates/profile", Profile)
	http.HandleFunc("/templates/register", Register)
	http.HandleFunc("/templates/logout", Logout)
	http.HandleFunc("/index.html", Home)
	http.HandleFunc("/", Home)
	http.HandleFunc("/templates/mkpost", Addposts)
	http.HandleFunc("/templates/post", Postfunc)
	http.HandleFunc("/temlates/admin", Admin)
	fmt.Printf("Started server successfully on http://localhost:1709/\n")
	http.ListenAndServe(":1709", nil)
}
