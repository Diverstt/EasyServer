package main

import (
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "You can sand only GET request", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("Hello, World!"))
}

func Start() {
	http.HandleFunc("/", GetHandler)
	http.ListenAndServe(":8080", nil)
}

func main() {
	Start()
}
