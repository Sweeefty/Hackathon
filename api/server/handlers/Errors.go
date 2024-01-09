package handlers

import (
	"Hackathon/api/server/data"
	"Hackathon/api/tools"
	"net/http"
)

func Error404(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("404 called")
	w.Header().Set("Content-Type", "application/json")

	Response := data.Response{
		Status: "Bad gateway :)",
		Code:   "404",
	}
	tools.ResponseF(w, Response)
}
