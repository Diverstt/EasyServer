package main

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Запрос: %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", LoggerMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Write([]byte("Middleware Test"))
		} else {
			http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		}
	})))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
