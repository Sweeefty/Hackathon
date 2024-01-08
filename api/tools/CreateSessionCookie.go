package tools

// We create a random string of 32 characters

import (
	"crypto/rand"
	"database/sql"
	"io"

	_ "github.com/mattn/go-sqlite3"
)

func CreateSessionCookie(email string) string {
	randomBytes := make([]byte, 32)
	io.ReadFull(rand.Reader, randomBytes)
	// We check if the cookie already exists
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		return "Error opening database"
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM accounts WHERE SessionCookie = ?", randomBytes)
	if err != nil {
		return "ErrorCheckingCookie"
	}
	defer rows.Close()
	if rows.Next() {
		// If it exists, we create a new one
		return CreateSessionCookie(email)
	} else {
		// If it doesn't exist, we return the cookie
		// Update the cookie in the database
		stmt, err := db.Prepare("UPDATE accounts SET SessionCookie = ? WHERE email = ?")
		if err != nil {
			return "Error updating cookie"
		}
		defer stmt.Close()
		_, err = stmt.Exec(randomBytes, email)
		if err != nil {
			return "Error updating cookie"
		}
		return string(randomBytes)
	}
}
