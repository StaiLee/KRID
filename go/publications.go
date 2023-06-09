package module

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Getuser(w http.ResponseWriter, r *http.Request) { //informations sur L'utilisateur
	db, _ := sql.Open("sqlite3", "./database.db")
	u, _ := r.Cookie("uuid")
	if err := db.QueryRow("SELECT creationdate, username, level from users where uuid = ?", u.Value).Scan(&user.Creationdate, &user.Username, &user.Level); err != nil { //request going to the database
		return
	}
}

func Getposts(w http.ResponseWriter, r *http.Request) { //affichages des 10 derniers post sur le site
	db, _ := sql.Open("sqlite3", "./database.db")
	rows, _ := db.Query("SELECT * FROM posts LIMIT 10")
	defer rows.Close()
	templ.Post = nil
	for rows.Next() {
		if err := rows.Scan(&post.Id, &post.Creator, &post.Title, &post.Content, &post.Picture, &post.Likes, &post.Dislikes, &post.Slash, &post.Date); err != nil {
			return
		}
		templ.Post = append(templ.Post, post)
	}
	defer db.Close()
}

func Getpostid(w http.ResponseWriter, r *http.Request) { //afficher un post specifique
	db, _ := sql.Open("sqlite3", "./database.db")
	post_id := r.URL.Query()["post-id"]
	rows, _ := db.Query("SELECT * FROM posts WHERE postid = ? LIMIT 3", post_id[0])
	defer rows.Close()
	templ.Post = nil
	for rows.Next() {
		if err := rows.Scan(&post.Id, &post.Creator, &post.Title, &post.Content, &post.Picture, &post.Likes, &post.Dislikes, &post.Slash, &post.Date); err != nil {
			return
		}
		templ.Post = append(templ.Post, post)
	}
	defer db.Close()
}

func Addposts(w http.ResponseWriter, r *http.Request) { //ajouter un poste sur le forum
	if !Ifregistered(w, r) {
		return
	}
	_ = r.ParseForm()

	picture := uploadHandler(w, r)
	if picture == "" {
		picture = "none"
	}
	title := r.PostForm.Get("title")
	if title == "" {
		template.Must(template.ParseFiles(filepath.Join(templatesDir, "./templates/mkpost.html"))).Execute(w, templ)
		return
	}
	creator, _ := r.Cookie("uuid")
	strcreator := creator.String()
	strcreator = strcreator[5:]
	content := r.PostForm.Get("content")
	slash := r.PostForm.Get("slash")
	dt := time.Now()
	dt.Format("01-02-2006 15:04:05")

	db, _ := sql.Open("sqlite3", "./database.db")
	request, _ := db.Prepare("INSERT INTO posts(creator, title, content, picture, likes, dislikes, slash, date) VALUES (?, ?, ?, ?, ?, ?, ?, ?)") // Preparation de la requete
	_, _ = request.Exec(strcreator, title, content, picture, 0, 0, slash, dt.String())
	http.Redirect(w, r, "/index.html", http.StatusFound)
	defer request.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request) string { //uploader un fichier
	r.Body = http.MaxBytesReader(w, r.Body, 2*1024*1024) // 2 Mb
	if err := r.ParseMultipartForm(2 << 20); err != nil {
		fmt.Print(err)
		return ""
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		return "no picture"
	}
	// buffer pour handle le fichier
	fileHeader := make([]byte, 512)
	if _, err := file.Read(fileHeader); err != nil {
		return ""
	}

	if _, err := file.Seek(0, 0); err != nil {
		return ""
	}
	filetype := http.DetectContentType(fileHeader)
	if filetype == "image/jpeg" || filetype == "image/gif" || filetype == "image/png" || filetype == "image/jpg" {
		out, pathError := ioutil.TempFile("./static/assets/temp-images", "*")
		if pathError != nil {
			log.Println("Error Creating a file for writing", pathError)
			return ""
		}
		defer out.Close()

		_, copyError := io.Copy(out, file)
		if copyError != nil {
			log.Println("Error copying", copyError)
		}
		str := out.Name()
		str = str[28:] //donne a la db le bon chemin vers le fichier
		return str
	}
	return ""
}

func Postfunc(w http.ResponseWriter, r *http.Request) { //main func
	Ratelimit(w, r)
	Ifregistered(w, r)
	Addcomment(w, r)
	Getpostid(w, r)
	Getcomments(w, r)
	template.Must(template.ParseFiles(filepath.Join(templatesDir, "./templates/post.html"))).Execute(w, templ)
}
