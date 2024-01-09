package handlers

import (
	"Hackathon/api/server/data"
	"Hackathon/api/tools"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("Root called")
	w.Header().Set("Content-Type", "application/json")
	Response := data.Response{
		Status: "ok",
		Code:   "200",
		Data:   map[string][]string{"endpoints": {"/connection"}},
	}
	tools.ResponseF(w, Response)
}
