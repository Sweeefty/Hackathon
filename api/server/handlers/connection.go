package handlers

import (
	"Hackathon/api/server/data"
	"Hackathon/api/tools"
	"net/http"
)

func ConnectionHandler(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("Connection called")
	w.Header().Set("Content-Type", "application/json")
	// email and password
	if r.Method == "GET" {
		email := r.FormValue("email")
		password := r.FormValue("password")
		// check if email and password are correct
		Equal := tools.Connection(email, password)
		if Equal != "Error" {
			// create a session cookie
			Cookie := tools.CreateSessionCookie(Equal)
			Response := data.Response{
				Status: "ok",
				Code:   "200",
				Data:   map[string]string{"cookie": Cookie},
			}
			tools.ResponseF(w, Response)

		} else {
			Response := data.Response{
				Status: "email or password incorrect",
				Code:   "400",
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
