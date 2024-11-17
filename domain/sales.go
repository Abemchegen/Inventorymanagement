package domain

import "time"

type Sales struct {
	ID            string      `json:"id" bson:"_id"`                        // Unique identifier
	Date          time.Time   `json:"date" bson:"date"`                     // Date of the transaction
	CustomerID    string      `json:"customer_id" bson:"customer_id"`       // Reference to the customer
	Items         []SalesItem `json:"items" bson:"items"`                   // List of items sold
	TotalAmount   float64     `json:"total_amount" bson:"total_amount"`     // Total monetary amount
	PaymentMethod string      `json:"payment_method" bson:"payment_method"` // Payment method
}

type SalesItem struct {
	ProductID string  `json:"product_id" bson:"product_id"` // ID of the product sold
	Quantity  int     `json:"quantity" bson:"quantity"`     // Quantity sold
	Price     float64 `json:"price" bson:"price"`           // Unit price of the product
}
