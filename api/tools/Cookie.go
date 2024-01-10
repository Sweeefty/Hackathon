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

func CreateSessionCookie(id string) string {
	random, err := RandomString(30)
	// We check if the cookie already exists
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		return "Error opening database"
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM session WHERE SessionCookie = ?", random)
	if err != nil {
		return "ErrorCheckingCookie"
	}
	defer rows.Close()
	if rows.Next() {
		// If it exists, we create a new one
		return CreateSessionCookie(id)
	} else {
		// If it doesn't exist, we return the cookie
		// Update the cookie in the database
		_, err = db.Exec("INSERT INTO session (accounts_id,SessionCookie) VALUES (?,?)", id, random)
		if err != nil {
			WriteErr("Error inserting into database tools/Inventory.go line 64")
			return "Error inserting into database"
		}
		return random
	}
}

//We check if the cookie is in the database
func CookieCheck(cookie string) bool {
	db, _ := sql.Open("sqlite3", "./data/db.sqlite")
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM session WHERE SessionCookie = ?", cookie)
	defer rows.Close()
	for rows.Next() {
		return true
	}
	return false
}
