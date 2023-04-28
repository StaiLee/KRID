package module

import (
	"fmt"
	"net/http"
)

func Start() {
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("./static/assets/css/"))))
	http.Handle("/assets/js/", http.StripPrefix("/assets/js/", http.FileServer(http.Dir("./static/assets/js/"))))
	http.Handle("/assets/images/", http.StripPrefix("/assets/images/", http.FileServer(http.Dir("./static/assets/images/"))))
	fmt.Printf("Started server successfully on http://localhost:1709/\n")
	http.ListenAndServe(":1709", nil)
}
