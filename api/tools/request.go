package tools

import (
	"Hackathon/api/server/data"
	"database/sql"
)

func GetRequestByCampus(id string) (data.RequestList, error) {
	var requestList data.RequestList
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/request.go line 11")
		return requestList, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id , accounts_id , campus_id , comment , title , anonymous FROM request WHERE campus_id = ?  AND read = false;", id)
	defer rows.Close()
	if err != nil {
		WriteErr("Error querying database tools/request.go line 31")
		return requestList, err
	}
	for rows.Next() {
		var request data.Request
		rows.Scan(&request.Id, &request.Account_id, &request.Campus_id, &request.Comment, &request.Title, &request.Anonymous)
		requestList.RequestList = append(requestList.RequestList, request)
	}
	return requestList, nil
}

func GetRequestById(id string) (data.Request, error) {
	var request data.Request
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/request.go line 11")
		return request, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id , accounts_id , campus_id , comment , title , anonymous FROM request WHERE id = ?;", id)
	defer rows.Close()
	if err != nil {
		WriteErr("Error querying database tools/request.go line 31")
		return request, err
	}
	for rows.Next() {
		rows.Scan(&request.Id, &request.Account_id, &request.Campus_id, &request.Comment, &request.Title, &request.Anonymous)
		return request, nil
	}
	return request, err
}

func CreateRequest(bde_id string, campus_id string, comment string, title string, anonymous string) error {
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/request.go line 11")
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO request (accounts_id, campus_id, comment, title , anonymous) VALUES (?,?,?,?,?);", bde_id, campus_id, comment, title, anonymous)
	if err != nil {
		WriteErr("Error inserting into database tools/Inventory.go line 64")
		return err
	}
	return nil
}

func ReadRequest(id string) error {
	db, err := sql.Open("sqlite3", "./data/db.sqlite")
	if err != nil {
		WriteErr("Error opening database tools/request.go line 11")
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE request SET read = (true) WHERE id = ?;", id)
	if err != nil {
		WriteErr("Error inserting into database tools/Inventory.go line 64")
		return err
	}
	return nil
}
