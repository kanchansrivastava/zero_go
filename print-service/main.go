package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("!!!!!!!!! Pimggggg")
	response := "This is my first Service and it works !!"
	w.Write([]byte(response)) // Convert string to []byte and write to ResponseWriter
}

func main() {
	http.HandleFunc("/health-check", pingHandler)
	panic(http.ListenAndServe(":80", nil))
}
