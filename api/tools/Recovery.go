package tools

import (
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
	// we check the recovery code
	if r.Method == "GET" {
		password := r.FormValue("psw")
		// check if email and password are correct
		if password == output {
			// we create response
			ResponseF(w, "200", output, map[string]string{"Code erreur système": output})
		} else {
			ResponseF(w, "400", "password incorrect", nil)
		}
	}
}
