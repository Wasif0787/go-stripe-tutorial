// cors.go
package main

import (
	"net/http"
)

// CORS middleware to enable CORS for all origins
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins for testing
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if r.Method == http.MethodOptions {
			return // Preflight request, respond and return
		}
		next.ServeHTTP(w, r)
	})
}
