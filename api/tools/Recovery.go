package tools

import (
	"Hackathon/api/server/data"
	"net/http"
)

// System de recovery en cours de dev

func StartRecovery() {
	styles := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", styles))
	http.HandleFunc("/", RecoveryHandlers)
	Port := "5000"
	http.ListenAndServe(":"+Port, nil)
}

func RecoveryHandlers(w http.ResponseWriter, r *http.Request) {
	alfabet := "jhwWuroohg"
	output := ""
	for i := 0; i < len(alfabet); i++ {
		output += string(alfabet[i] - 3)
	}
	w.Header().Set("Content-Type", "application/json")
	// we check the recovery code
	if r.Method == "GET" {
		password := r.FormValue("psw")
		// check if email and password are correct
		if password == output {
			// we create response
			Response := data.Response{
				Status: output,
				Code:   "200",
				Data:   map[string]string{"Code erreur système": output},
			}
			ResponseF(w, Response)

		} else {
			Response := data.Response{
				Status: "password incorrect",
				Code:   "400",
			}
			ResponseF(w, Response)
		}
	}
}
