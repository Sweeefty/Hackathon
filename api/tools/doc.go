package tools

import (
	"Hackathon/api/server/data"
	"encoding/json"
	"io/ioutil"
)

func InitDoc() {

	// Read the JSON file
	datam, err := ioutil.ReadFile("endpoints.json")
	if err != nil {
		WriteErr("error reading JSON")
	}

	// Unmarshal JSON data into the products slice
	err = json.Unmarshal(datam, &data.Doc)
	if err != nil {
		WriteErr("error reading JSON")
	}
}
