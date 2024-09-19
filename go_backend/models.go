package main

// HealthResponse represents the structure of the health check response.
type HealthResponse struct {
	Status  string `json:"status"`  // The status of the server (e.g., "OK").
	Message string `json:"message"` // A message indicating the server's state.
}
