package module

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Dellike(w http.ResponseWriter, r *http.Request) { //enleve les likes
	print("idem")
}

func Deldislike(w http.ResponseWriter, r *http.Request) { //enleves les dislike
	print("idem")
}

func Adddislike(w http.ResponseWriter, r *http.Request) { //ajout des likes dans un post
	_ = r.ParseForm()
	db, _ := sql.Open("sqlite3", "./database.db")
	id := r.URL.Query()["dislike"]
	if id == nil {
		return
	}
	postid := id[0]
	var dislikes int
	if err := db.QueryRow("SELECT dislikes from posts where postid = ?", postid).Scan(&dislikes); err != nil { //requete vers la database
		return
	}
	dislikes += 1
	_, _ = db.Exec("UPDATE posts SET dislikes = ? WHERE postid = ?", dislikes, postid)
}

func Addlike(w http.ResponseWriter, r *http.Request) { //ajout d'un like dans un post
	if !Ifregistered(w, r) {
		return
	}
	_ = r.ParseForm()
	db, _ := sql.Open("sqlite3", "./database.db")
	id := r.URL.Query()["like"]
	if id == nil {
		return
	}
	postid := id[0]
	var likes int
	u, err := r.Cookie("uuid")
	uuid := u.String()
	if err != nil {
		return
	}
	uuid = uuid[5:]
	if err := db.QueryRow("SELECT karma FROM interaction WHERE postid = ? AND uuid = ?", postid, uuid).Scan(&likes); err != sql.ErrNoRows && err != nil { //request going to the database
		return
	}
	if likes == 1 { //si l'utilisateur a deja like
		return
	}
	dt := time.Now()
	dt.Format("01-02-2006 15:04:05")

	db, _ = sql.Open("sqlite3", "./database.db")
	reqdata := "INSERT INTO interaction(postid, uuid, karma, date) VALUES (?, ?, ?, ?)"
	request, _ := db.Prepare(reqdata) // prepare la reuete
	_, err = request.Exec(postid, uuid, 1, dt)
	fmt.Print(err)

	defer request.Close()
	if err := db.QueryRow("SELECT likes from posts where postid = ?", postid).Scan(&likes); err != nil { //request va dans la database
		fmt.Print(err)
		return
	}
	likes += 1
	_, err = db.Exec("UPDATE posts SET likes = ? WHERE postid = ?", likes, postid)
	fmt.Print(err)
	db.Close()
}

func Addcomment(w http.ResponseWriter, r *http.Request) { //ajout d'un commentaire dans un post
	if !Ifregistered(w, r) {
		return
	}
	_ = r.ParseForm()
	db, _ := sql.Open("sqlite3", "./database.db")
	file := uploadHandler(w, r)
	comment := r.PostForm.Get("comment")
	if comment == "" {
		return
	}

	post_id := r.URL.Query()["post-id"]
	u, _ := r.Cookie("uuid")
	uuid := u.String()
	uuid = uuid[5:]
	var Com Comment
	if err := db.QueryRow("SELECT creationdate, username, level from users where uuid = ?", uuid).Scan(&Com.Creationdateuser, &Com.Username, &Com.Level); err != nil { //requete Database
		return
	}
	dt := time.Now()
	dt.Format("01-02-2006 15:04:05")

	reqdata := "INSERT INTO interaction(postid, uuid, karma, date) VALUES (?, ?, ?, ?)"
	request, _ := db.Prepare(reqdata) // Prepare la requete
	_, _ = request.Exec(post_id, uuid, 1, dt)
	defer request.Close()

	reqdata = "INSERT INTO comments(creator, postid, comment, likes, dislikes, file, creationdateuser, username, level, date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	request, _ = db.Prepare(reqdata) // Prepare la requete
	_, _ = request.Exec(uuid, post_id[0], comment, 0, 0, file, Com.Creationdateuser, Com.Username, Com.Level, dt)
	defer request.Close()
}

func Getcomments(w http.ResponseWriter, r *http.Request) { //tout les commentaires d'un post
	db, _ := sql.Open("sqlite3", "./database.db")
	post_id := r.URL.Query()["post-id"]
	postint, _ := strconv.Atoi(post_id[0])
	rows, _ := db.Query("SELECT * FROM comments WHERE postid = ?", &postint)
	var Com Comment //seul commentaire
	templ.Comments = nil
	for rows.Next() {
		if err := rows.Scan(&Com.Creator, &Com.Postid, &Com.Comment, &Com.Likes, &Com.Dislikes, &Com.File, &Com.Creationdateuser, &Com.Username, &Com.Level, &Com.Date); err != nil {
			return
		}
		templ.Comments = append(templ.Comments, Com)
	}
	defer rows.Close()
}
