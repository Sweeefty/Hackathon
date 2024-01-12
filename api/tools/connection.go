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
	rows, err := db.Query("SELECT id FROM accounts WHERE email = ? AND password = ?", email, password)
	defer rows.Close()
	var id string
	if rows.Next() {
		rows.Scan(&id)
		return id
	} else {
		return "Error"
	}
}
