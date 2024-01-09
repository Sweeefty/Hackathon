package tools

import (
	"database/sql"
)

func GetInventory(cookie string) interface{} {
	// open database
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/Inventory.go line 11")
		return "Error opening database"
	}
	defer db.Close()

	//Extract id from cookie
	rows, err := db.Query("SELECT id FROM accounts WHERE SessionCookie = ?", cookie)
	defer rows.Close()
	if err != nil {
		WriteErr("Error querying database tools/Inventory.go line 20")
		return "Error querying database"
	}
	var id int
	for rows.Next() {
		rows.Scan(&id)
	}
	// Extract inventory from id
	rows, err = db.Query("SELECT products_id FROM inventory WHERE accounts_id = ?", id)
	defer rows.Close()
	if err != nil {
		WriteErr("Error querying database tools/Inventory.go line 31")
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

func AddInventory(cookie string, ObjectId string) bool {
	// open database
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/Inventory.go line 54")
		return false
	}
	defer db.Close()

	//Extract id from cookie
	rows, err := db.Query("SELECT id FROM accounts WHERE SessionCookie = ?", cookie)
	defer rows.Close()
	if err != nil {
		WriteErr("Error querying database tools/Inventory.go line 54")
		return false
	}
	var id int
	for rows.Next() {
		rows.Scan(&id)
	}

	_, err = db.Exec("INSERT INTO inventory (accounts_id, products_id) VALUES (?, ?)", id, ObjectId)
	if err != nil {
		WriteErr("Error inserting into database tools/Inventory.go line 64")
		return false
	}
	return true
}
