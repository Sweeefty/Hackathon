package handlers

import (
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
		} else if authHeader != "test" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized: Invalid Authorization header")
			return
		}
		next.ServeHTTP(w, r)
	})
}
