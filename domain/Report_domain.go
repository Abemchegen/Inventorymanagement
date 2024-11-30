package domain

type Report struct {
	BestSellingProduct []BestSellingProduct `json:"best_selling_product" bson:"best_selling_product"`
	Revenue            int                  `json:"revenue" bson:"revenue"`
	Profit             int                  `json:"profit" bson:"profit"`
}

type BestSellingProduct struct {
	ProductID         string `json:"product_id" bson:"product_id"`
	ProductName       string `json:"product_name" bson:"product_name"`
	Category          string `json:"category" bson:"category"`
	RemainingQuantity int    `json:"remaining_quantity" bson:"remaining_quantity"`
	TurnOver          int    `json:"turn_over" bson:"turn_over"`
	Price             int    `json:"price" bson:"price"`
}

type BestSellingCategory struct {
	Category  string `json:"category" bson:"category"`
	TurnOver  int    `json:"turn_over" bson:"turn_over"`
	ProductID string `json:"product_id" bson:"product_id"`
}

type OverView struct {
	TotalProduct  int `json:"total_product" bson:"total_product"`
	TotalRevenue  int `json:"total_revenue" bson:"total_revenue"`
	TotalProfit   int `json:"total_profit" bson:"total_profit"`
	TotalCost     int `json:"total_cost" bson:"total_cost"`
	TotalSold     int `json:"total_sold" bson:"total_sold"`
	TotalCategory int `json:"total_category" bson:"total_category"`
	TotalSupplier int `json:"total_supplier" bson:"total_supplier"`
	TopSelling    int `json:"top_selling" bson:"top_selling"`
	TotalOrders   int `json:"total_orders" bson:"total_orders"`
}

type ReportRepository interface {
	GetBestSellingProduct() (Report, error)
	GetBestSellingCategory() ([]BestSellingCategory, error)
	GetOverView() (OverView, error)
}

type ReportUseCase interface {
	GetBestSellingProduct() (Report, error)
	GetBestSellingCategory() ([]BestSellingCategory, error)
	GetOverView() (OverView, error)
}