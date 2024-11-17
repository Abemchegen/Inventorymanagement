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

type SupplierUseCase interface {
	GetSupplierByID(id string) (*Supplier, error)
	CreateSupplier(supplier *Supplier) error
	GetAllSuppliers() ([]Supplier, error)
}

type SupplierRepository interface {
	FindByID(id string) (*Supplier, error)
	FindAll() ([]Supplier, error)
	Store(supplier *Supplier) error
}
