package domain

type Product struct {
	ID          string  `json:"id" bson:"_id"`                  // Unique identifier (MongoDB ObjectId)
	Name        string  `json:"name" bson:"name"`               // Name of the product
	Description string  `json:"description" bson:"description"` // Detailed description
	Quantity    int     `json:"quantity" bson:"quantity"`       // Current stock quantity
	UnitPrice   float64 `json:"unit_price" bson:"unit_price"`   // Price per unit
	Category    string  `json:"category" bson:"category"`       // Product category
	SupplierID  string  `json:"supplier_id" bson:"supplier_id"` // Reference to the supplier
	Status      string  `json:"status" bson:"status"`           // Current status (e.g., in stock, sold)
}

type ProductUsecase interface {
	GetProductByID(id string) (*Product, error)
	GetAllProducts() ([]Product, error)
	CreateProduct(product *Product) error
	UpdateProduct(id string, product *Product) error
	DeleteProduct(id string) error
}

type ProductRepository interface {
	FindByID(id string) (*Product, error)
	FindAll() ([]Product, error)
	Store(product *Product) error
	Update(id string, product *Product) error
	Delete(id string) error
}
