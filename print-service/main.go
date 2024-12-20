package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("!!!!!!!!! Pimggggg")
	response := "Print Service is up !! After the CI/CD"
	w.Write([]byte(response)) // Convert string to []byte and write to ResponseWriter
}

func main() {
	http.HandleFunc("/health-check", pingHandler)
	panic(http.ListenAndServe(":80", nil))
}
