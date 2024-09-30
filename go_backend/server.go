package main

import (
	"log"
	"net/http"
)

// main initializes the HTTP server, maps routes to handlers, and starts listening for requests.
// It runs on port 4242, handling both payment intent creation and health check routes.
func main() {
	// Map the /create-payment-intent route to handleCreatePaymentIntent function.
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)

	// Map the /health route to handleHealth function for health check purposes.
	http.HandleFunc("/health", handleHealth)

	// Log that the server is starting and listen on port 4242.
	log.Println("Server is running on port 4242")

	// Start the HTTP server and handle any errors during startup.
	if err := http.ListenAndServe(":4242", nil); err != nil {
		log.Fatalf("ERROR: Failed to start server: %v", err)
	}
}
