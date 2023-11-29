package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
type  data struct {
data map[string]interface{}
}
*/

var data map[string]interface{}

func main() {
	d := "https://ads.pubmatic.com/AdServer/js/pwt/floors/158451/3767/floors.json"
	s, e := http.Get(d)
	if e == nil {
		println()
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(s.Body)
	respBytes := buf.String()

	err := json.Unmarshal([]byte(respBytes), &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	} else {

		for _, interest := range (data["modelGroups"]).([]interface{}) {
			jsonData, err := json.Marshal(interest)
			if err != nil {
				fmt.Println("Error encoding JSON:", err)
				return
			}

			// Printing JSON data
			fmt.Println("Encoded JSON:", string(jsonData))

		}

	}
}
