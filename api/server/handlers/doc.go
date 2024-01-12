package handlers

import (
	"Hackathon/api/server/data"
	"Hackathon/api/tools"
	"net/http"
)

func Doc(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("doc called")
	if r.Method == "GET" {
		tools.ResponseF(w, "200", "request successful", data.Doc)
	} else {
		tools.ResponseF(w, "400", "method incorrect", nil)
	}
}
