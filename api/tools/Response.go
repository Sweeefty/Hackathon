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
		fmt.Println(err)
	}
	w.Write([]byte(jsonResp))
	//json.NewEncoder(w).Encode(Response)
}
