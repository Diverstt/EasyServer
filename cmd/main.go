package main

import (
	"fmt"
	"log"
	"net/http"
)

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "You can send only GET request\n")
			return
		} else {
			log.Printf("%s %s\n", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		}
	})
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Middleware Test\n")
}

func startServer() {
	http.Handle("/", loggingMiddleware(mainHandler))

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	startServer()
}
