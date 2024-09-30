package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/paymentintent"
)

// handleCreatePaymentIntent processes HTTP POST requests to create Stripe payment intents.
// It expects a JSON body with order details and returns the client secret for the payment intent.
func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {
	// Load environment variables from the .env file.
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatalf("Error loading .env file: %v", envErr)
	}

	// Ensure the request method is POST, else return a 405 Method Not Allowed response.
	if request.Method != http.MethodPost {
		log.Printf("WARN: Invalid method %s on /create-payment-intent", request.Method)
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Decode the incoming request body into the Request struct.
	var req Request
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		log.Printf("ERROR: Failed to decode request body: %v", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Stripe secret key from environment variables.
	stripe.Key = os.Getenv("STRIPE_SECRET")

	// Create payment intent parameters with order amount and automatic payment methods.
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(calculateOrdersAmount(req.ProductID)), // Calculate amount based on product
		Currency: stripe.String(string(stripe.CurrencyUSD)),          // Use USD as the currency
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	// Create a new Stripe payment intent and handle any errors.
	paymentIntent, err := paymentintent.New(params)
	if err != nil {
		log.Printf("ERROR: Failed to create payment intent: %v", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare the response containing the client secret for the payment intent.
	var res Response
	res.ClientSecret = paymentIntent.ClientSecret
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(res)
	if err != nil {
		log.Printf("ERROR: Failed to encode response: %v", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set content type to JSON and send the response back to the client.
	writer.Header().Set("Content-Type", "application/json")
	if _, err = io.Copy(writer, &buf); err != nil {
		log.Printf("ERROR: Failed to write response: %v", err)
	}
}

// handleHealth serves the /health endpoint, providing the status of the server.
// Responds with a 200 OK and a JSON message when the server is healthy.
func handleHealth(writer http.ResponseWriter, request *http.Request) {
	// Prepare the health response with status and message.
	response := HealthResponse{
		Status:  "OK",
		Message: "Server is up and running!",
	}

	// Set response content type to application/json and return 200 OK.
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	// Encode the response to JSON and send it back to the client.
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		log.Printf("ERROR: Failed to write health response: %v", err)
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}
