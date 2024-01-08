package tools

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Connection(email string, password string) string {
	// open database
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		return "Error opening database"
	}
	defer db.Close()

	//Check if the password is correct
	rows, err := db.Query("SELECT * FROM accounts WHERE email = ? AND password = ?", email, password)
	defer rows.Close()
	if rows.Next() {
		return "ok"
	} else {
		return "Error"
	}
}
