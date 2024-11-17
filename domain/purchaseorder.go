package domain

import "time"

type PurchaseOrder struct {
	ID                   string    `json:"id" bson:"_id"`                                        // Unique identifier
	SupplierID           string    `json:"supplier_id" bson:"supplier_id"`                       // Reference to the supplier
	OrderDate            time.Time `json:"order_date" bson:"order_date"`                         // Date when the order was created
	ExpectedDeliveryDate time.Time `json:"expected_delivery_date" bson:"expected_delivery_date"` // Anticipated delivery date
	Items                []POItem  `json:"items" bson:"items"`                                   // List of items ordered
	Status               string    `json:"status" bson:"status"`                                 // Order status (e.g., pending, completed)
}

type POItem struct {
	ProductID string `json:"product_id" bson:"product_id"` // ID of the product
	Quantity  int    `json:"quantity" bson:"quantity"`     // Quantity ordered
}
