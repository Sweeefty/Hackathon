package tools

import (
	"database/sql"
)

func GetInventory(cookie string) interface{} {
	// open database
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		return "Error opening database"
	}
	defer db.Close()

	//Extract id from cookie
	rows, err := db.Query("SELECT id FROM accounts WHERE SessionCookie = ?", cookie)
	defer rows.Close()
	if err != nil {
		return "Error querying database"
	}
	var id int
	for rows.Next() {
		rows.Scan(&id)
	}
	// Extract inventory from id
	rows, err = db.Query("SELECT * FROM inventory WHERE id = ?", id)
	defer rows.Close()
	if err != nil {
		return "Error querying database"
	}
	var inventory []string
	for rows.Next() {
		var item string
		rows.Scan(&item)
		inventory = append(inventory, item)
	}
	return inventory
}
