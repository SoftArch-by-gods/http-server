package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	const RFC3339 = "2006-01-02T15:04:05Z07:00"
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now().Format(RFC3339)
		fmt.Println(now)
	})
	http.ListenAndServe(":8795", nil)
}