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

func Register(w http.ResponseWriter, r *http.Request) { //connexion a un compte
	Ifregistered(w, r)
	_ = r.ParseForm()
	db, _ := sql.Open("sqlite3", "./database.db")
	username := r.PostForm.Get("identifiant")
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	date := time.Now().Format("02-01-2006")
	if username == "" { //execution de la template
		template.Must(template.ParseFiles(filepath.Join(templatesDir, "./templates/register.html"))).Execute(w, templ)
		return
	}
	u, _ := uuid.NewV4()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	reqdata := `INSERT INTO users(uuid, username, email, password, creationdate, level, karma) VALUES (?, ?, ?, ?, ?, ?, ?)`
	request, _ := db.Prepare(reqdata) // preparation de la requete
	_, _ = request.Exec(u.String(), username, email, hashedPassword, date, 1, 0)
	defer request.Close()
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "uuid", Value: u.String(), Expires: expiration, Path: "/"}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/index.html", http.StatusFound)
}

func Login(w http.ResponseWriter, r *http.Request) { //connexion login a un compte existant
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
	if err := db.QueryRow("SELECT password from users where email = ?", email).Scan(&passworddb); err != nil { //requete allant dans la database
		template.Must(template.ParseFiles(filepath.Join(templatesDir, "./templates/login.html"))).Execute(w, templ)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(passworddb), []byte(password)); err != nil { //comparaison entre le mot de passe, et la version de ce dernier hasher
		w.WriteHeader(http.StatusUnauthorized) // mot de passe incorrect
	}
	//mot de passe correct
	var uuid string
	if err := db.QueryRow("SELECT uuid from users where email = ?", email).Scan(&uuid); err != nil { //requete allant dans la database
		template.Must(template.ParseFiles(filepath.Join(templatesDir, "./templates/login.html"))).Execute(w, templ)
		return
	}
	expiration := time.Now().Add(365 * 24 * time.Hour) //prise des uuid et mise en place des cookies
	cookie := http.Cookie{Name: "uuid", Value: uuid, Expires: expiration, Path: "/"}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/index.html", http.StatusFound)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	u, _ := r.Cookie("uuid")
	cookie := http.Cookie{Name: "uuid", Value: u.String(), Expires: time.Unix(0, 0), Path: "/"}
	http.SetCookie(w, &cookie)
	user.Level = 0
	templ.User = user
	http.Redirect(w, r, "/index.html", http.StatusFound)
	template.Must(template.ParseFiles(filepath.Join(templatesDir, "./static/index.html"))).Execute(w, templ)
}

func Users(w http.ResponseWriter, r *http.Request) {
	print("hmmm")
}

func Promoteuser(w http.ResponseWriter, r *http.Request) {
	print("hmmm")
}

func Ifregistered(w http.ResponseWriter, r *http.Request) bool {
	_, err := r.Cookie("uuid")
	if err != nil {
		user.Level = 0
		return false
	}
	Getuser(w, r)
	templ.User = user
	return true
}
