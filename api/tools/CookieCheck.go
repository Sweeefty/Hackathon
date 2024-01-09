package tools

// We create a random string of 32 characters

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//We check if the cookie is in the database
func CookieCheck(cookie string) bool {
	db, _ := sql.Open("sqlite3", "./data/db.sqlite")
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM accounts WHERE SessionCookie = ?", cookie)
	defer rows.Close()
	for rows.Next() {
		return true
	}
	return false
}
