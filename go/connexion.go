package module

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) { //register into your account
	Ifregistered(w, r)
	_ = r.ParseForm()
	db, _ := sql.Open("sqlite3", "./database.db")
	username := r.PostForm.Get("identifiant")
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	date := time.Now().Format("02-01-2006")
	if username == "" { //first load of page
		template.Must(template.ParseFiles(filepath.Join(templatesDir, "./templates/register.html"))).Execute(w, templ)
		return
	}
	u, _ := uuid.NewV4()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	reqdata := `INSERT INTO users(uuid, username, email, password, creationdate, level, karma) VALUES (?, ?, ?, ?, ?, ?, ?)`
	request, _ := db.Prepare(reqdata) // Prepare request.
	_, _ = request.Exec(u.String(), username, email, hashedPassword, date, 1, 0)
	defer request.Close()
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "uuid", Value: u.String(), Expires: expiration, Path: "/"}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/index.html", http.StatusFound)
}

func Login(w http.ResponseWriter, r *http.Request) { //login to your account
	Ifregistered(w, r)
	_ = r.ParseForm()
	db, _ := sql.Open("sqlite3", "./database.db")
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	var passworddb string
	if email == "" {
		template.Must(template.ParseFiles(filepath.Join(templatesDir, "./templates/login.html"))).Execute(w, templ)
		return
	}
	if err := db.QueryRow("SELECT password from users where email = ?", email).Scan(&passworddb); err != nil { //request going to the database
		template.Must(template.ParseFiles(filepath.Join(templatesDir, "./templates/login.html"))).Execute(w, templ)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(passworddb), []byte(password)); err != nil { //comparison between hashed password and input password
		w.WriteHeader(http.StatusUnauthorized) // error passwords do not match
	}
	//password matches
	var uuid string
	if err := db.QueryRow("SELECT uuid from users where email = ?", email).Scan(&uuid); err != nil { //request going to the database
		template.Must(template.ParseFiles(filepath.Join(templatesDir, "./templates/login.html"))).Execute(w, templ)
		return
	}
	expiration := time.Now().Add(365 * 24 * time.Hour) //get uuid from database and set cookie with it
	cookie := http.Cookie{Name: "uuid", Value: uuid, Expires: expiration, Path: "/"}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/index.html", http.StatusFound)
}

func Logout(w http.ResponseWriter, r *http.Request) { //remove cookie on logout
	u, _ := r.Cookie("uuid")
	cookie := http.Cookie{Name: "uuid", Value: u.String(), Expires: time.Unix(0, 0), Path: "/"}
	http.SetCookie(w, &cookie)
	user.Level = 0
	templ.User = user
	http.Redirect(w, r, "/index.html", http.StatusFound)
	template.Must(template.ParseFiles(filepath.Join(templatesDir, "./static/index.html"))).Execute(w, templ)
}

func Users(w http.ResponseWriter, r *http.Request) { //get the users of the site
	print("hmmm")
}

func Promoteuser(w http.ResponseWriter, r *http.Request) { //promote a user
	print("hmmm")
}

func Ifregistered(w http.ResponseWriter, r *http.Request) bool { //is the user registered ?
	_, err := r.Cookie("uuid")
	if err != nil {
		user.Level = 0
		return false
	}
	Getuser(w, r)
	templ.User = user
	return true
}
