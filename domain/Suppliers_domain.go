package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Suppliers struct {
	ID            primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Name          string             `json:"name" bson:"name"`
	Product       string             `json:"product" bson:"product"`
	ContactNumber string             `json:"contact_number" bson:"contact_number"`
	Email         string             `json:"email" bson:"email"`
	Image         string             `json:"image" bson:"image"`
	Category      string             `json:"category" bson:"category"`
}

type SuppliersUseCase interface {
	CreateSuppliers(suppliers Suppliers) (Suppliers, error)
	GetAllSuppliers() ([]Suppliers, error)
	GetSuppliersByID(id string) (Suppliers, error)
	DeleteSuppliers(id string) (Suppliers,error)
	UpdateSuppliers(suppliers Suppliers) (Suppliers,error)
}
type SuppliersRepository interface {
	CreateSuppliers(suppliers Suppliers) (Suppliers, error)
	GetAllSuppliers() ([]Suppliers, error)
	GetSuppliersByID(id string) (Suppliers, error)
	DeleteSuppliers(id string) (Suppliers,error)
	UpdateSuppliers(suppliers Suppliers) (Suppliers,error)

}