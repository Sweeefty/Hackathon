package handlers

import (
	"Hackathon/api/server/data"
	"Hackathon/api/tools"
	"fmt"
	"net/http"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			Response := data.Response{
				Status: "",
				Code:   "200",
			}
			tools.ResponseF(w, Response)
			return
		} else if authHeader != "test" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized: Invalid Authorization header")
			return
		}
		next.ServeHTTP(w, r)
	})
}
