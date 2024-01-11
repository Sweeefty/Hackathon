package handlers

import (
	"Hackathon/api/tools"
	"net/http"
)

func GetInventory(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("GetInventory called")
	if r.Method == "GET" {
		cookie := r.FormValue("cookie")
		if tools.CookieCheck(cookie) {
			Inventory := tools.GetInventory(cookie)
			tools.ResponseF(w, "200", "request successful", Inventory)

		} else {
			tools.ResponseF(w, "400", "cookie incorrect", nil)

		}
	} else {
		tools.ResponseF(w, "400", "method incorrect", nil)
	}
}

func AddInventory(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("AddInventory called")
	if r.Method == "POST" {
		cookie := r.FormValue("cookie")
		ObjectId := r.FormValue("ObjectId")
		if tools.CookieCheck(cookie) {
			if tools.AddInventory(cookie, ObjectId) {
				tools.ResponseF(w, "200", "request successful", nil)
			} else {
				tools.ResponseF(w, "400", "internal error please retry later", nil)
			}

		} else {
			tools.ResponseF(w, "400", "cookie incorrect", nil)

		}
	} else {
		tools.ResponseF(w, "400", "method incorrect", nil)
	}
}
