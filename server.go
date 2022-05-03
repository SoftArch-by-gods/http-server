package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Date struct {
	Time string `json:"time"`
}

func main() {
	const RFC3339 = "2006-01-02T15:04:05Z07:00"

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now().Format(RFC3339)
		inStruct := Date{now}

		data, err := json.MarshalIndent(inStruct, "", "  ")
		if err != nil {
			panic(err)
		}
		result := string(data)
		fmt.Fprintf(w, result)
	})
	http.ListenAndServe(":8795", nil)
}
