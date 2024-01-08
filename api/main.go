package main

import (
	"Hackathon/api/server"
	"Hackathon/api/tools"
)

func main() {
	// We create the database if it doesn't exist
	tools.CreateDB()
	// We start the API
	server.StartApi()
}
