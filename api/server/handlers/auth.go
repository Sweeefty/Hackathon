package handlers

import (
	"Hackathon/api/tools"
	"fmt"
	"net/http"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized: Missing Authorization header")
			return
		} else if !tools.CheckAuthorization(authHeader) {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized: Invalid Authorization header")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func SetAuthorization(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("SetAuthorization called")

	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")
		// check if email and password are correct
		Equal := tools.Connection(email, password)
		if Equal != "Error" {
			Account, err := tools.GetAccountById(Equal)
			if Account.Role == "admin" && err == nil {
				token := tools.CreateAuthorization()
				tools.WriteLog("SetAuthorization created. Admin account: " + email)
				tools.ResponseF(w, "200", "success", map[string]string{"Token": token})
			} else {
				tools.ResponseF(w, "400", "The account is not admin", nil)
			}
		} else {
			tools.ResponseF(w, "400", "email or password incorrect", nil)
		}

	} else {
		tools.ResponseF(w, "400", "method incorrect", nil)
	}
}
