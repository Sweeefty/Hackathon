package handlers

import (
	"Hackathon/api/server/data"
	"Hackathon/api/tools"
	"net/http"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("GetEvents called")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		Events := tools.GetEvents("")

		if Events {
			Response := data.Response{
				Status: "ok",
				Code:   "200",
				Data:   Events,
			}
			tools.ResponseF(w, Response)
		} else {
			Response := data.Response{
				Status: "events not found",
				Code:   "404",
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

func AddEvent(w http.ResponseWriter, r *http.Request) {

	tools.WriteLog("AddEvent called")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		//Add the event
		tools.AddEvent("")
		//Response
		Response := data.Response{
			Status: "ok",
			Code:   "200",
		}
		tools.ResponseF(w, Response)
	} else {
		Response := data.Response{
			Status: "method incorrect",
			Code:   "400",
		}
		tools.ResponseF(w, Response)
	}
}