package handlers

import (
	"Hackathon/api/server/data"
	"Hackathon/api/tools"
	"net/http"
)

func GetInventory(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	// email and password
	if r.Method == "GET" {
		cookie := r.FormValue("cookie")
		if true {
			Inventory := tools.GetInventory(cookie)
			Response := data.Response{
				Status: "ok",
				Code:   "200",
				Data:   Inventory,
			}
			tools.Response(w, Response)

		} else {
			Response := data.Response{
				Status: "cookie incorrect",
				Code:   "400",
			}
			tools.Response(w, Response)
		}
	}
}
