package module

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func Home(w http.ResponseWriter, r *http.Request) {
	Ratelimit(w, r)
	Ifregistered(w, r)
	Getposts(w, r)
	Addlike(w, r)
	Notifications(w, r)
	template.Must(template.ParseFiles(filepath.Join(templatesDir, "./static/index.html"))).Execute(w, templ)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	Ratelimit(w, r)
	if Ifregistered(w, r) {
		Getprofileinfo(w, r)
		template.Must(template.ParseFiles(filepath.Join(templatesDir, "./templates/profile.html"))).Execute(w, templ)
	}
}

func Aboutus(w http.ResponseWriter, r *http.Request) {
	Ratelimit(w, r)
	Ifregistered(w, r)
	var admin jsonUser // api github fetch
	if templ.JsonUser == nil {
		urls := [2]string{"https://api.github.com/users/luxchar", "https://api.github.com/users/naywvi"}
		for _, url := range urls {
			// request http api
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}

			// read body
			body, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}

			if res.StatusCode != 200 {
				log.Fatal("Unexpected status code", res.StatusCode)
			}

			err = json.Unmarshal(body, &admin)
			if err != nil {
				log.Fatal(err)
			}
			templ.JsonUser = append(templ.JsonUser, admin)
		}
	}
	template.Must(template.ParseFiles(filepath.Join(templatesDir, "./templates/about.html"))).Execute(w, templ)
}

func Admin(w http.ResponseWriter, r *http.Request) { //remove cookie on logout
	Ratelimit(w, r)
	db, _ := sql.Open("sqlite3", "./database.db")
	if !Ifregistered(w, r) {
		return
	}
	uuid, _ := r.Cookie("uuid")
	u := uuid.String()
	u = u[5:]
	level := 0
	if err := db.QueryRow("SELECT level from users where uuid = ?", u).Scan(&level); err != nil { //request going to the database
		return
	}
	if level == 3 { //check user level
		template.Must(template.ParseFiles(filepath.Join(templatesDir, "./templates/admin.html"))).Execute(w, templ)
	}
}
