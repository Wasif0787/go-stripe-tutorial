package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Register the handler function for the /create-payment-intent endpoint.
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)

	// Log the server startup message.
	log.Println("Server is running on port 4242")

	// Start the HTTP server on port 4242.
	// If an error occurs, log it and terminate the application.
	err := http.ListenAndServe(":4242", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// handleCreatePaymentIntent processes requests to the /create-payment-intent endpoint.
func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	// Log a message to indicate that this endpoint has been hit.
	fmt.Println("Create Payment Intent function called")
}
