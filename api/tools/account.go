package tools

import (
	"Hackathon/api/server/data"
	"database/sql"
)

func GetAccountById(id string) interface{} {
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
