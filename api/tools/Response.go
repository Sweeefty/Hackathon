package tools

import (
	"Hackathon/api/server/data"
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseF(w http.ResponseWriter, Response data.Response) {
	jsonResp, err := json.Marshal(Response)
	if err != nil {
		WriteErr("Error marshalling response tools/Response.go l13")
		fmt.Println(err)
	}
	w.Write([]byte(jsonResp))
}
