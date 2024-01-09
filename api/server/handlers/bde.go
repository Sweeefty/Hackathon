package handlers

import (
	"Hackathon/api/server/data"
	"Hackathon/api/tools"
	"net/http"
)

func GetBde(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("GetProducts called")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		bde_id := r.FormValue("id")
		if bde_id == "" {
			Products := tools.GetBde()
			if Products != false {
				Response := data.Response{
					Status: "request successful",
					Code:   "200",
					Data:   Products,
				}
				tools.ResponseF(w, Response)
			} else {
				Response := data.Response{
					Status: "error",
					Code:   "400",
				}
				tools.ResponseF(w, Response)
			}
		} else {
			Products := tools.GetBdeById(bde_id)
			if Products != false {
				Response := data.Response{
					Status: "request successful",
					Code:   "200",
					Data:   Products,
				}
				tools.ResponseF(w, Response)
			} else {
				Response := data.Response{
					Status: "error",
					Code:   "400",
				}
				tools.ResponseF(w, Response)

			}
		}

	} else {
		Response := data.Response{
			Status: "method incorrect",
			Code:   "400",
		}
		tools.ResponseF(w, Response)
	}
}
