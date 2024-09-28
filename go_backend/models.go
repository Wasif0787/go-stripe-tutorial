package main

// HealthResponse represents the structure of the health check response.
type HealthResponse struct {
	Status  string `json:"status"`  // The status of the server (e.g., "OK").
	Message string `json:"message"` // A message indicating the server's state.
}

type Request struct {
	ProductID string `json:"product_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
}
