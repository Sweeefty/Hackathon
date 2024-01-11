package handlers

import (
	"Hackathon/api/tools"
	"net/http"
)

func GetRequest(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("GetRequest called")

	if r.Method == "GET" {
		bde_id := r.FormValue("bdeId")
		campus_id := r.FormValue("campusId")
		request_id := r.FormValue("requestId")
		if bde_id != "" {
			campus_id, err := tools.GetCampusIdByBdeId(bde_id)
			if err != nil {
				tools.ResponseF(w, "400", "Incorrect bde id", nil)
				return
			}
			Requests, err := tools.GetRequestByCampus(campus_id)
			if err != nil {
				tools.ResponseF(w, "500", "internal server error", nil)
				return
			}
			tools.ResponseF(w, "200", "success", Requests)
		} else if campus_id != "" {
			Requests, err := tools.GetRequestByCampus(campus_id)
			if err != nil {
				tools.ResponseF(w, "500", "internal server error", nil)
				return
			}
			tools.ResponseF(w, "200", "success", Requests)
		} else if request_id != "" {
			Request, err := tools.GetRequestById(request_id)
			if err != nil {
				tools.ResponseF(w, "400", "Incorrect id", nil)
				return
			}
			tools.ResponseF(w, "200", "success", Request)
		} else {
			tools.ResponseF(w, "400", "missing parameter(s)", nil)
		}
	} else {
		tools.ResponseF(w, "400", "method incorrect", nil)
	}
}

func CreateRequest(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("CreateRequest called")

	if r.Method == "POST" {
		bde_id := r.FormValue("bdeId")
		comment := r.FormValue("comment")
		title := r.FormValue("title")
		anonymous := r.FormValue("anonymous")
		if bde_id != "" || comment != "" || title != "" || anonymous != "" {
			campus_id, err := tools.GetCampusIdByBdeId(bde_id)
			if err != nil {
				tools.ResponseF(w, "400", "Incorrect bde id", nil)
				return
			}
			err = tools.CreateRequest(bde_id, campus_id, comment, title, anonymous)
			if err != nil {
				tools.ResponseF(w, "500", "internal server error", nil)
				return
			}
			tools.ResponseF(w, "200", "success", nil)
		} else {
			tools.ResponseF(w, "400", "missing parameter(s)", nil)
		}
	} else {
		tools.ResponseF(w, "400", "method incorrect", nil)
	}
}
