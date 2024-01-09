package handlers

import (
	"Hackathon/api/server/data"
	"Hackathon/api/tools"
	"net/http"
)

func GetInventory(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("GetInventory called")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		cookie := r.FormValue("cookie")
		if tools.CookieCheck(cookie) {
			Inventory := tools.GetInventory(cookie)
			Response := data.Response{
				Status: "ok",
				Code:   "200",
				Data:   Inventory,
			}
			tools.ResponseF(w, Response)

		} else {
			Response := data.Response{
				Status: "cookie incorrect",
				Code:   "400",
			}
			tools.ResponseF(w, Response)
		}
	}
}

func AddInventory(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("AddInventory called")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		cookie := r.FormValue("cookie")
		ObjectId := r.FormValue("ObjectId")
		if tools.CookieCheck(cookie) {
			if tools.AddInventory(cookie, ObjectId) {
				Response := data.Response{
					Status: "ok",
					Code:   "200",
				}
				tools.ResponseF(w, Response)
			} else {
				Response := data.Response{
					Status: "error",
					Code:   "500",
				}
				tools.ResponseF(w, Response)
			}

		} else {
			Response := data.Response{
				Status: "cookie incorrect",
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
