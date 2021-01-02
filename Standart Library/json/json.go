package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `
		{
			"data": {
				"object":"card",
				"id":"card_545645454",
				"first_name":"Onur",
				"last_name":"Teber",
				"balance":"955.537"
			}
		}
	`

	var m map[string]map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		panic(err)
	}

	fmt.Println(m)

	fmt.Println("**********************")

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

}
