package tools

import (
	"Hackathon/api/server/data"
	"database/sql"
)

func GetAccountById(id string) (data.Account, error) {
	var Account data.Account
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/Account.go line 11")
		return Account, err
	}
	defer db.Close()

	// Extract inventory from id
	rows, err := db.Query("SELECT id , email , role , campus_id , bde_id FROM accounts WHERE id = ?;", id)
	defer rows.Close()
	if err != nil {
		WriteErr("Error querying database tools/Account.go line 31")
		return Account, err
	}
	for rows.Next() {
		rows.Scan(&Account.Id, &Account.Email, &Account.Role, &Account.Campus_id, &Account.Bde_id)
		return Account, nil
	}
	return Account, err
}

func GetAccountByCookie(cookie string) interface{} {
	id := GetIdByCookie(cookie)
	// Open the database
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/Account.go line 11")
		return false
	}
	defer db.Close()

	// Extract inventory from id
	rows, err := db.Query("SELECT id , email , role , campus_id , bde_id FROM accounts WHERE id = ?;", id)
	defer rows.Close()
	if err != nil {
		WriteErr("Error querying database tools/Account.go line 31")
		return false
	}
	for rows.Next() {
		var Account data.Account
		rows.Scan(&Account.Id, &Account.Email, &Account.Role, &Account.Campus_id, &Account.Bde_id)
		return Account
	}
	return false
}

func GetIdByCookie(cookie string) string {
	// open database
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/Inventory.go line 11")
		return "Error opening database"
	}
	defer db.Close()

	//Extract id from cookie
	rows, err := db.Query("SELECT id FROM session WHERE SessionCookie = ?", cookie)
	defer rows.Close()
	if err != nil {
		WriteErr("Error querying database tools/Inventory.go line 20")
		return "Error querying database"
	}
	var id string
	for rows.Next() {
		rows.Scan(&id)
	}
	return id
}

func GetBdeIdById(id string) interface{} {
	// Open database
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/Inventory.go line 11")
		return false
	}
	defer db.Close()

	// Extract id from cookie
	rows, err := db.Query("SELECT bde_id FROM accounts WHERE bde_id = ?", id)
	defer rows.Close()
	if err != nil {
		WriteErr("Error querying database tools/Inventory.go line 20")
		return "Error querying database"
	}

	for rows.Next() {
		rows.Scan(&id)
	}
	return id
}
