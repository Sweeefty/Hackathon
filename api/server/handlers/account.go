package handlers

import (
	"Hackathon/api/server/data"
	"Hackathon/api/tools"
	"net/http"
)

func GetAccountInfo(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("Connection called")
	w.Header().Set("Content-Type", "application/json")
	// email and password
	if r.Method == "GET" {
		id := r.FormValue("id")
		Account := tools.GetAccountById(id)
		if Account != false {
			Response := data.Response{
				Status: "ok",
				Code:   "200",
				Data:   Account,
			}
			tools.ResponseF(w, Response)
		} else {
			Response := data.Response{
				Status: "account not found",
				Code:   "404",
			}
			tools.ResponseF(w, Response)
		}

	} else {
		Response := data.Response{
			Status: "method incorrect",
			Code:   "400",
		}
		tools.ResponseF(w, Response)
	}
}
