package tools

// We create a random string of 32 characters

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"

	_ "github.com/mattn/go-sqlite3"
)

func RandomString(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	randomString := base64.URLEncoding.EncodeToString(randomBytes)
	return randomString[:length], nil
}

func CreateSessionCookie(email string) string {
	random, err := RandomString(30)
	// We check if the cookie already exists
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		return "Error opening database"
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM accounts WHERE SessionCookie = ?", random)
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
		_, err = stmt.Exec(random, email)
		if err != nil {
			return "Error updating cookie"
		}
		return random
	}
}
