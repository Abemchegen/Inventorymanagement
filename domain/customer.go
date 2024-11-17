package domain

type Customer struct {
	ID            string `json:"id" bson:"_id"`                        // Unique identifier
	Name          string `json:"name" bson:"name"`                     // Customer's full name
	ContactNumber string `json:"contact_number" bson:"contact_number"` // Phone number
	Email         string `json:"email" bson:"email"`                   // Email address
	Address       string `json:"address" bson:"address"`               // Physical address
	LoyaltyPoints int    `json:"loyalty_points" bson:"loyalty_points"` // Points for customer loyalty
}
