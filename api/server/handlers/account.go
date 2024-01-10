package handlers

import (
	"Hackathon/api/tools"
	"net/http"
)

func GetAccountInfo(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("Connection called")
	// email and password
	if r.Method == "GET" {
		cookie := r.FormValue("cookie")
		if cookie == "" {
			tools.ResponseF(w, "400", "missing parameter (cookie)", nil)
			return
		}

		Account := tools.GetAccountByCookie(cookie)
		if Account != false {
			tools.ResponseF(w, "200", "Successful", Account)
		} else {
			tools.ResponseF(w, "403", "cookie incorrect , not found", nil)

		}

	} else {
		tools.ResponseF(w, "400", "method incorrect", nil)
	}
}
