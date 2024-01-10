package handlers

import (
	"Hackathon/api/tools"
	"net/http"
)

func ConnectionHandler(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("Connection called")
	// email and password
	if r.Method == "GET" {
		email := r.FormValue("email")
		password := r.FormValue("password")
		// check if email and password are correct
		Equal := tools.Connection(email, password)
		if Equal != "Error" {
			// create a session cookie
			Cookie := tools.CreateSessionCookie(Equal)
			tools.ResponseF(w, "200", "request successful", map[string]string{"cookie": Cookie})

		} else {
			tools.ResponseF(w, "400", "email or password incorrect", nil)
		}
	} else {
		tools.ResponseF(w, "400", "method incorrect", nil)
	}
}
