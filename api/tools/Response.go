package tools

import (
	"Hackathon/api/server/data"
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseF(w http.ResponseWriter, code string, Status string, Data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	var Response interface{}
	if Data == nil {
		Response = data.ResponseDataLess{
			Status: Status,
			Code:   code,
		}
	} else {
		Response = data.Response{
			Status: Status,
			Code:   code,
			Data:   Data,
		}
	}

	jsonResp, err := json.Marshal(Response)
	if err != nil {
		WriteErr("Error marshalling response tools/Response.go l13")
		fmt.Println(err)
	}
	w.Write([]byte(jsonResp))
}
