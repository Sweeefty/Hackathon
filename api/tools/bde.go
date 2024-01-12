package tools

import (
	"Hackathon/api/server/data"
	"database/sql"
)

func GetBde() interface{} {
	// open database
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/bde.go line 11")
		return false
	}
	defer db.Close()

	// Extract inventory from id
	rows, err := db.Query("SELECT id , name ,campus_id FROM bde;")
	defer rows.Close()
	if err != nil {
		WriteErr("Error querying database tools/bde.go line 31")
		return false
	}
	BdeList := data.BdeList{}
	for rows.Next() {
		var Bde data.Bde
		rows.Scan(&Bde.Name, &Bde.Id, &Bde.Campus_id)
		BdeList.Bde = append(BdeList.Bde, Bde)
	}
	return BdeList
}

func GetBdeById(id interface{}) interface{} {
	// open database
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/Inventory.go line 11")
		return false
	}
	defer db.Close()

	// Extract bde from id
	rows, err := db.Query("SELECT id , name ,campus_id FROM bde WHERE id = ?;", id)
	defer rows.Close()
	if err != nil {
		WriteErr("Error querying database tools/bde.go line 31")
		return false
	}
	for rows.Next() {
		var Bde data.Bde
		rows.Scan(&Bde.Name, &Bde.Id, &Bde.Campus_id)
		return Bde
	}
	return false
}
