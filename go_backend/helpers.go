package main

// calculateOrdersAmount calculates the total order amount based on the provided product ID.
// It returns the price in cents (e.g., $26.00 -> 2600 cents).
func calculateOrdersAmount(productID string) int64 {
	// Pricing logic based on product ID. Returns price in cents.
	switch productID {
	case "Forever Pants":
		return 26000 // $260.00
	case "Forever Shirt":
		return 15500 // $155.00
	case "Forever Shorts":
		return 30000 // $300.00
	}
	// Default amount if no valid product ID is provided.
	return 0
}
