package handlers

import (
	"Hackathon/api/tools"
	"net/http"
)

func Error404(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("404 called")
	tools.ResponseF(w, "404", "Bad gateway :)", nil)
}
