package domain

type Supplier struct {
	ID            string `json:"id" bson:"_id"`                                  // Unique identifier
	Name          string `json:"name" bson:"name"`                               // Supplier name
	ContactPerson string `json:"contact_person" bson:"contact_person"`           // Contact person's name
	ContactNumber string `json:"contact_number" bson:"contact_number"`           // Contact phone number
	Email         string `json:"email" bson:"email"`                             // Contact email
	Address       string `json:"address" bson:"address"`                         // Supplier's address
	PricingInfo   string `json:"pricing_information" bson:"pricing_information"` // Pricing details
}
