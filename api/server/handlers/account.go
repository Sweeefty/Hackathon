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
		cookie := r.FormValue("cookie")
		if cookie == "" {
			Response := data.Response{
				Status: "missing parameter (cookie)",
				Code:   "400",
			}
			tools.ResponseF(w, Response)
			return
		}
		Account := tools.GetAccountByCookie(cookie)
		if Account != false {
			Response := data.Response{
				Status: "ok",
				Code:   "200",
				Data:   Account,
			}
			tools.ResponseF(w, Response)
		} else {
			Response := data.Response{
				Status: "cookie incorrect , not found",
				Code:   "403",
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
