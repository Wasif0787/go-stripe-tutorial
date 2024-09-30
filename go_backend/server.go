package main

import (
	"crypto/tls"
	"encoding/base64"
	"log"
	"net/http"
	"os"
)

func main() {
	// Read the certificate from the environment variable
	certBase64 := os.Getenv("CERT_BASE64")
	if certBase64 == "" {
		log.Fatalf("ERROR: CERT_BASE64 environment variable is not set")
	}
	certPEM, err := base64.StdEncoding.DecodeString(certBase64)
	if err != nil {
		log.Fatalf("ERROR: Failed to decode cert: %v", err)
	}

	// Read the key from the environment variable
	keyBase64 := os.Getenv("KEY_BASE64")
	if keyBase64 == "" {
		log.Fatalf("ERROR: KEY_BASE64 environment variable is not set")
	}
	keyPEM, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		log.Fatalf("ERROR: Failed to decode key: %v", err)
	}

	// Load the certificates into a tls.Certificate
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		log.Fatalf("ERROR: Failed to load X509 key pair: %v", err)
	}

	// Create a custom HTTP server with the TLS configuration
	server := &http.Server{
		Addr:    ":443",
		Handler: enableCors(http.DefaultServeMux),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	// Map the routes
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/health", handleHealth)

	// Log that the server is starting and listen on port 443
	log.Println("Server is running on port 443")

	// Start the HTTPS server
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("ERROR: Failed to start server: %v", err)
	}
}
