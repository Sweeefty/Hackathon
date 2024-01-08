package tools

// We create a random string of 32 characters

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CookieCheck(cookie string) bool {
	db, _ := sql.Open("sqlite3", "./data/db.sqlite")
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM accounts WHERE SessionCookie = ?", cookie)
	defer rows.Close()
	if rows.Next() {
		return true
	} else {
		return false
	}
}
