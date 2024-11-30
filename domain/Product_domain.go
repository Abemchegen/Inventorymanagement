package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID primitive.ObjectID `bson:"_id,omitempity" json:"id"`
	ProductID string `bson:"product_id" json:"product_id"`
	Name string `bson:"name" json:"name"`
	BuyPrice int `bson:"buy_price" json:"buy_price"`
	Quantity  int `bson:"quantity" json:"quantity"`
	ThresholdValue int `bson:"threshold_value" json:"threshold_value"`
	ExpiryDate string `bson:"expiry_date" json:"expiry_date"`
	Avaliable bool `bson:"avaliable" json:"avaliable"`
	Category string `bson:"category" json:"category"`
	Unit string `bson:"unit" json:"unit"`
	HowManySold int `bson:"how_many_sold" json:"how_many_sold" default:"0"`
	SellPrice int `bson:"sell_price" json:"sell_price" default:"0"`
}

type ProductUseCase interface {
	CreateProduct(product Product) (Product, error)
	GetAllProduct() ([]Product, error)
	GetProductByID(id string) (Product, error)
	UpdateProduct(product Product) (Product, error)
	DeleteProduct(id string) (Product, error)
}

type ProductRepository interface {
	CreateProduct(product Product) (Product, error)
	GetAllProduct() ([]Product, error)
	GetProductByID(id string) (Product, error)
	UpdateProduct(product Product) (Product, error)
	DeleteProduct(id string) (Product, error)
}
