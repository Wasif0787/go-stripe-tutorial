package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// handleCreatePaymentIntent processes HTTP requests to the /create-payment-intent endpoint.
// It expects a POST request and handles the creation of a payment intent.
func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {
	// Check if the request method is POST.
	if request.Method != http.MethodPost {
		// Log a warning for invalid method requests.
		log.Printf("WARN: Invalid method %s on /create-payment-intent", request.Method)
		// Respond with 405 Method Not Allowed.
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Log the information that the payment intent function has been called.
	log.Println("INFO: Create Payment Intent function called")
	// Respond with a 200 OK status and a success message.
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Payment Intent created successfully"))
}

// handleHealth responds to health check requests at the /health endpoint.
// It returns the server status in JSON format.
func handleHealth(writer http.ResponseWriter, request *http.Request) {
	// Create a HealthResponse object to hold the status information.
	response := HealthResponse{
		Status:  "OK",
		Message: "Server is up and running!",
	}

	// Set the Content-Type header to application/json for the response.
	writer.Header().Set("Content-Type", "application/json")
	// Respond with a 200 OK status.
	writer.WriteHeader(http.StatusOK)

	// Encode the response object as JSON and write it to the response writer.
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		// Log an error message if the response writing fails.
		log.Printf("ERROR: Failed to write response: %v", err)
		// Respond with a 500 Internal Server Error if an error occurs.
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}
