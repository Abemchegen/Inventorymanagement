package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Orders struct {
	ID          primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	ProductName string             `json:"product_name" bson:"product_name"`
	OrderValue int                `json:"order_value" bson:"order_value"`
	Quantity   int                `json:"quantity" bson:"quantity"`
	ExpectedDelivery string       `json:"expected_delivery" bson:"expected_delivery"`
	Status string                `json:"status" bson:"status"`
	ProductID string             `json:"product_id" bson:"product_id"`
	CreatedAt time.Time             `json:"created_at" bson:"created_at"`
	DateOFdelivery string        `json:"date_of_delivery" bson:"date_of_delivery"`

}

type OrdersUseCase interface {
	CreateOrders(orders Orders) ( Orders , error)
	GetAllOrders() ([]Orders, error)
	UpdateOrders(orders Orders) ( Orders , error)
	DeleteOrders(id string) (Orders, error)
	GetOrdersByID(id string) (Orders, error)
}
type OrdersRepository interface {
	CreateOrders(orders Orders) ( Orders , error)
	GetAllOrders() ([]Orders, error)
	UpdateOrders(orders Orders) ( Orders , error)
	DeleteOrders(id string) (Orders, error)
	GetOrdersByID(id string) (Orders, error)
}

