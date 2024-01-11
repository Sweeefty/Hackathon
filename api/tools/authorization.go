package tools

import (
	"database/sql"
	"time"
)

func CreateAuthorization() string {
	token, _ := RandomString(20)
	// open database
	db, err := sql.Open("sqlite3", "./data/dbAuth.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/Inventory.go line 54")
		return "false"
	}
	defer db.Close()
	hash, err := HashPassword(token)
	if err != nil {
		WriteErr("Error hashing password tools/Inventory.go line 59")
		return "false"
	}

	currentDate := time.Now().Format("2006-01-02")
	_, err = db.Exec("INSERT INTO authorization (hash_key, date) VALUES (?, ?);", hash, currentDate)
	if err != nil {
		WriteErr("Error inserting into database tools/Inventory.go line 64")
		return "false"
	}
	return token
}

func CheckAuthorization(token string) bool {
	// open database
	db, err := sql.Open("sqlite3", "./data/dbAuth.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/Inventory.go line 54")
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT hash_key FROM authorization;")
	defer rows.Close()
	if err != nil {
		WriteErr("Error querying database tools/Inventory.go line 20")
		return false
	}

	for rows.Next() {
		var hash string
		rows.Scan(&hash)
		if CheckPasswordHash(token, hash) {
			return true
		}
	}
	return false
}
