package module

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func Setup() {
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("./static/assets/css/"))))
	http.Handle("/assets/js/", http.StripPrefix("/assets/js/", http.FileServer(http.Dir("./static/assets/js/"))))
	http.Handle("/assets/images/", http.StripPrefix("/assets/images/", http.FileServer(http.Dir("./static/assets/images/"))))
	http.Handle("/assets/img/", http.StripPrefix("/assets/img/", http.FileServer(http.Dir("./static/assets/img/"))))
	http.Handle("/assets/temp-images/", http.StripPrefix("/assets/temp-images/", http.FileServer(http.Dir("./static/assets/temp-images/"))))
	http.HandleFunc("/pages/login", Login)
	http.HandleFunc("/pages/about", Aboutus)
	http.HandleFunc("/pages/profile", Profile)
	http.HandleFunc("/pages/register", Register)
	http.HandleFunc("/pages/logout", Logout)
	http.HandleFunc("/index.html", Home)
	http.HandleFunc("/", Home)
	http.HandleFunc("/pages/mkpost", Addposts)
	http.HandleFunc("/pages/post", Postfunc)
	http.HandleFunc("/pages/admin", Admin)
	fmt.Printf("Started server successfully on http://localhost:1709/\n")
	http.ListenAndServe(":1709", nil)
}
