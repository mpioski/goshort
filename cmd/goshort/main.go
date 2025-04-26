package main

import (
	"log"
	"net/http"

	"github.com/mpioski/goshort/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", handler.Shorten)
	mux.HandleFunc("GET /", handler.Redirect)

	port := "8080"
	log.Printf("Starting server on port %s", port)
	err := http.ListenAndServe(":"+port, loggingMiddleware(mux))
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
