package tools

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetEvents(event string) bool {
	db, _ := sql.Open("sqlite3", "./data/query/db.sqlite")
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM events WHERE name = ?", event)
	defer rows.Close()
	for rows.Next() {
		return true
	}
	return false
}

func AddEvent(event string) {
	db, _ := sql.Open("sqlite3", "./data/query/db.sqlite")
	defer db.Close()
	db.Exec("INSERT INTO events (name) VALUES (?)", event)
}