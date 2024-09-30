package main

// HealthResponse represents the JSON response for the /health endpoint.
type HealthResponse struct {
	Status  string `json:"status"`  // The status of the server (e.g., "OK").
	Message string `json:"message"` // A message indicating the server's state.
}

// Request represents the incoming JSON request to create a payment intent.
type Request struct {
	ProductID string `json:"product_id"` // The product identifier for the order.
	FirstName string `json:"first_name"` // The customer's first name.
	LastName  string `json:"last_name"`  // The customer's last name.
	Email     string `json:"email"`      // The customer's email address.
	Address1  string `json:"address_1"`  // Primary address for billing.
	Address2  string `json:"address_2"`  // Secondary address for billing.
	City      string `json:"city"`       // The city for the billing address.
	State     string `json:"state"`      // The state for the billing address.
	Zip       string `json:"zip"`        // The postal code for the billing address.
	Country   string `json:"country"`    // The country for the billing address.
}

// Response represents the JSON response with the Stripe client secret.
type Response struct {
	ClientSecret string `json:"clientSecret"` // The Stripe payment intent's client secret.
}
