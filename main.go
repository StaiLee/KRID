package main

import (
	"database/sql"
	"log"
	Server "module/go"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	Server.Start()
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT,
        email TEXT,
		password TEXT
    );
`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
    INSERT INTO users (name, email) VALUES (?, ?);
`, "John Doe", "john@example.com")
	if err != nil {
		log.Fatal(err)
	}
}
