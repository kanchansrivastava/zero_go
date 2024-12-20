package main

import (
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	response := "This is my first Service and it works. Current version is 1 !!"
	w.Write([]byte(response)) // Convert string to []byte and write to ResponseWriter
}

func main() {
	http.HandleFunc("/health-check", pingHandler)
	panic(http.ListenAndServe(":80", nil))
}
