package module

import (
	"database/sql"
	"net/http"
)

func Reports(w http.ResponseWriter, r *http.Request) { //makes a report to admin
	_ = r.ParseForm()
	postid := r.URL.Query()["post-id"]
	if postid == nil {
		return
	}
	strpostid := postid[0]
	db, _ := sql.Open("sqlite3", "./database.db")
	report := r.PostForm.Get("content")
	uuid, _ := r.Cookie("uuid")
	u := uuid.String()
	u = u[5:]
	reqdata := "INSERT INTO report(report, user, postid) VALUES (?, ?, ?)"
	request, _ := db.Prepare(reqdata) // Prepare request.
	_, _ = request.Exec(report, u, strpostid)
	defer request.Close()
}

func Getreport(w http.ResponseWriter, r *http.Request) { //get the reports to see them
	var report Report
	db, _ := sql.Open("sqlite3", "./database.db")
	rows, _ := db.Query("SELECT report, user, postid, reportid FROM report LIMIT 10")
	defer rows.Close()
	templ.Report = nil
	for rows.Next() {
		if err := rows.Scan(&report.ReportContent, &report.User, &report.Postid, &report.ReportId); err != nil {
			return
		}
		templ.Report = append(templ.Report, report)
	}
}

func Getadmin(w http.ResponseWriter, r *http.Request) { //get the list of admins
	var mod Mod
	db, _ := sql.Open("sqlite3", "./database.db")
	rows, _ := db.Query("SELECT uuid, username, email, creationdate, level FROM users WHERE level = 2 OR level = 3")
	defer rows.Close()
	templ.Mods = nil
	for rows.Next() {
		if err := rows.Scan(&mod.Uuid, &mod.Username, &mod.Email, &mod.Creationdate, &mod.Level); err != nil {
			return
		}
		templ.Mods = append(templ.Mods, mod)
	}
}

func Delpost(w http.ResponseWriter, r *http.Request) { //remove post NOT TESTED
	postid := r.PostForm.Get("postid")
	db, _ := sql.Open("sqlite3", "./database.db")
	if err := db.QueryRow("DELETE FROM posts WHERE postid = ?", postid); err != nil { //request going to the database
		return
	}
	if err := db.QueryRow("DELETE FROM report WHERE post-id = ?", postid); err != nil { //request going to the database
		return
	}
}
