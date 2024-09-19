package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// HealthResponse represents the structure of the health check response.
type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	// Register the handler function for the /create-payment-intent endpoint.
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/health", handleHealth)

	// Log the server startup message.
	log.Println("INFO: Server is running on port 4242")

	// Start the HTTP server on port 4242.
	if err := http.ListenAndServe(":4242", nil); err != nil {
		log.Fatalf("ERROR: Failed to start server: %v", err)
	}
}

// handleCreatePaymentIntent processes requests to the /create-payment-intent endpoint.
func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		// Log and return a 405 Method Not Allowed if it's not a POST request.
		log.Printf("WARN: Invalid method %s on /create-payment-intent", request.Method)
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// TODO: Process payment logic here.
	log.Println("INFO: Create Payment Intent function called")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Payment Intent created successfully"))
}

// handleHealth responds with the server health status.
func handleHealth(writer http.ResponseWriter, request *http.Request) {
	// Log the health check request.
	log.Println("INFO: Health check endpoint called")

	// Prepare and send a JSON response.
	response := HealthResponse{
		Status:  "OK",
		Message: "Server is up and running!",
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(writer).Encode(response); err != nil {
		log.Printf("ERROR: Failed to write response: %v", err)
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}
