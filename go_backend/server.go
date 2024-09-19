package main

import (
	"log"
	"net/http"
)

// main is the entry point of the application. It sets up the HTTP server,
// registers route handlers, and starts listening on port 4242.
func main() {
	// Register the handler function for the /create-payment-intent endpoint.
	// This maps the route "/create-payment-intent" to the handleCreatePaymentIntent function.
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)

	// Register the handler function for the /health endpoint.
	// This maps the route "/health" to the handleHealth function.
	http.HandleFunc("/health", handleHealth)

	// Log a message to indicate that the server is starting and listening on port 4242.
	log.Println("Server is running on port 4242")

	// Start the HTTP server on port 4242. If an error occurs (e.g., port is already in use),
	// the server will log the error message and terminate the application.
	if err := http.ListenAndServe(":4242", nil); err != nil {
		log.Fatalf("ERROR: Failed to start server: %v", err)
	}
}
