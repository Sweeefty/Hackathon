package handlers

import (
	"Hackathon/api/server/data"
	"Hackathon/api/tools"
	"net/http"
)

func Doc(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("doc called")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		Response := data.Response{
			Status: "ok",
			Code:   "200",
			Data:   data.Doc,
		}
		tools.ResponseF(w, Response)
	} else {
		Response := data.Response{
			Status: "method incorrect",
			Code:   "400",
		}
		tools.ResponseF(w, Response)
	}
}