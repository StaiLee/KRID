package module

import (
	"database/sql"
	"fmt"
	"net/http"
)

func Getprofileinfo(w http.ResponseWriter, r *http.Request) { //gets notification for a user
	db, _ := sql.Open("sqlite3", "./database.db")
	creator, _ := r.Cookie("uuid")
	u := creator.String()
	u = u[5:]
	rows, _ := db.Query("SELECT * FROM posts WHERE creator = ? ORDER BY date LIMIT 3", u)
	templ.Post = nil
	var temppost Publication
	for rows.Next() {
		if err := rows.Scan(&temppost.Id, &temppost.Creator, &temppost.Title, &temppost.Content, &temppost.Picture, &temppost.Likes, &temppost.Dislikes, &temppost.Slash, &temppost.Date); err != nil {
			return
		}
		templ.Post = append(templ.Post, temppost)
	}
	defer rows.Close()

	rows, _ = db.Query("SELECT * FROM comments WHERE creator = ? ORDER BY date LIMIT 3", u)
	templ.Comments = nil
	var Com Comment
	for rows.Next() {
		if err := rows.Scan(&Com.Creator, &Com.Postid, &Com.Comment, &Com.Likes, &Com.Dislikes, &Com.File, &Com.Creationdateuser, &Com.Username, &Com.Level, &Com.Date); err != nil {
			return
		}
		templ.Comments = append(templ.Comments, Com)
	}
	defer rows.Close()

	rows, _ = db.Query("SELECT * FROM interaction WHERE uuid = ? ORDER BY date LIMIT 3", u)
	templ.Inter = nil
	var Inter Interaction
	for rows.Next() {
		if err := rows.Scan(&Inter.Postid, &Inter.Uuid); err != nil {
			return
		}
		templ.Inter = append(templ.Inter, Inter)
	}
	defer rows.Close()
}

func Notifications(w http.ResponseWriter, r *http.Request) { //gets notification for a user NOT YET FINISHED
	if !Ifregistered(w, r) {
		return
	}
	db, _ := sql.Open("sqlite3", "./database.db")
	creator, _ := r.Cookie("uuid")
	u := creator.String()
	u = u[5:]
	templ.Inter = nil
	rows, err := db.Query("SELECT * FROM posts WHERE creator = ? ORDER BY date", u) //fetch user posts
	fmt.Print(err)
	for rows.Next() {
		if err := rows.Scan(&Notif.Post.Id, &Notif.Post.Creator, &Notif.Post.Title, &Notif.Post.Content, &Notif.Post.Picture, &Notif.Post.Likes, &Notif.Post.Dislikes, &Notif.Post.Slash, &Notif.Post.Date); err != nil {
			fmt.Print(err)
			defer rows.Close()
			return
		}

		rows2, _ := db.Query("SELECT * FROM interaction WHERE postid = ? ORDER BY date LIMIT 2", &Notif.Post.Id) //fetch likes dropped under user posts
		var Inter Interaction
		for rows2.Next() {
			if err := rows.Scan(&Inter.Postid, &Inter.Uuid, &Inter.Date); err != nil {
				// fmt.Print(err)
				defer rows2.Close()
				defer rows.Close()
				return
			}
			templ.Inter = append(templ.Inter, Inter)
		}
		defer rows2.Close()

		rows3, _ := db.Query("SELECT * FROM comments WHERE postid = ? ORDER BY date LIMIT 3", &Notif.Post.Id) //fetch comments dropped under user posts
		for rows3.Next() {
			templ.Comments = nil
			var Com Comment
			if err := rows3.Scan(&Com.Creator, &Com.Postid, &Com.Comment, &Com.Likes, &Com.Dislikes, &Com.File, &Com.Creationdateuser, &Com.Username, &Com.Level, &Com.Date); err != nil {
				// fmt.Print(err)
				defer rows3.Close()
				defer rows.Close()
				return
			}
			templ.Notif.Com = append(templ.Notif.Com, Com)
		}
		defer rows3.Close()

		if len(templ.Notif.Com)+len(templ.Notif.Inter) >= 4 { //limit number of notifications to 4-5 max
			defer rows.Close()
			return
		}
	}
	defer rows.Close()
}
