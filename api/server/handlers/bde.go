package handlers

import (
	"Hackathon/api/tools"
	"net/http"
)

func GetBde(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("GetBde called")
	if r.Method == "GET" {
		bde_id := r.FormValue("id")
		if bde_id == "" {
			Products := tools.GetBde()
			if Products != false {
				tools.ResponseF(w, "200", "request successful", Products)
			} else {
				tools.ResponseF(w, "400", "internal error please retry later", nil)

			}
		} else {
			Products := tools.GetBdeById(bde_id)
			if Products != false {
				tools.ResponseF(w, "200", "request successful", Products)
			} else {
				tools.ResponseF(w, "400", "internal error please retry later", nil)

			}
		}

	} else {
		tools.ResponseF(w, "400", "method incorrect", nil)
	}
}

func AddBde(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		campus_id := r.FormValue("campus")
		if name != "" && campus_id != "" {
			Products := tools.GetBde()
			if Products != false {
				tools.ResponseF(w, "200", "request successful", Products)

			} else {
				tools.ResponseF(w, "400", "internal error please retry later", nil)
			}
		} else {
			tools.ResponseF(w, "400", "missing parameter(s)", nil)
		}

	} else {
		tools.ResponseF(w, "400", "method incorrect", nil)
	}
}
