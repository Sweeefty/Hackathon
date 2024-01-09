package main

import (
	"Hackathon/api/server"
	"Hackathon/api/tools"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "recover" || os.Args[1] == "r" {
			tools.StartRecovery()
		} else if os.Args[1] == "d" || os.Args[1] == "D" {
			tools.DeleteLog()
			tools.CreateLog()
		} else if os.Args[1] == "p" || os.Args[1] == "P" {
			tools.PrintLog()
		} else if os.Args[1] == "h" || os.Args[1] == "H" {
			tools.Help()
		} else {
			tools.Help()
		}
	} else {
		tools.WriteLog("Server started")
		// We create the database if it doesn't exist
		tools.CreateDB()
		// We initialize the documentation
		tools.InitDoc()
		// We start the API
		server.StartApi()
	}
}
