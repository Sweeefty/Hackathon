package tools

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDB() {
	dbFile := "./data/db.sqlite"
	// Check if the database file exists, if not, create it
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		file, err := os.Create(dbFile)
		if err != nil {
			fmt.Println("Error creating database file:", err)
			return
		}
		defer file.Close()
		// Open or create the SQLite database file
		db, err := sql.Open("sqlite3", dbFile)
		if err != nil {
			fmt.Println("Error opening database:", err)
			return
		}
		defer db.Close()

		// Read SQL query from a file
		queryFile := "data/query/createDB.sql"
		queryBytes, err := ioutil.ReadFile(queryFile)
		if err != nil {
			fmt.Println("Error reading query file:", err)
			return
		}
		query := string(queryBytes)

		// Execute the SQL query to create a table
		_, err = db.Exec(query)
		if err != nil {
			fmt.Println("Error executing query:", err)
			return
		}

		fmt.Println("Table created successfully!")
	}

}
