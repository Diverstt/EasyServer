package main

import (
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("Hello, World!"))
	}
}

func main() {
	http.HandleFunc("/", GetHandler)
	http.ListenAndServe(":8080", nil)
}
