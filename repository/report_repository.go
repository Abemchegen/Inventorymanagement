package repository

import (
	"context"
	"inventory/domain"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReportRepository struct {
	database   mongo.Database
	orderCollection string
	productCollection string
	sellerCollection string
	storeCollection string
}

func NewReportRepository(database mongo.Database, OrderCollection, ProductCollection, SellerCollection, StoreCollection string) domain.ReportRepository {
	return &ReportRepository{
		database:   database,
		orderCollection: OrderCollection,
		productCollection: ProductCollection,
		sellerCollection: SellerCollection,
		storeCollection: StoreCollection,
		}

}

func (a *ReportRepository) GetBestSellingProduct() (domain.Report, error) {
	var reports domain.Report
	var products []domain.Product
	var orders []domain.Orders
	var Suppliers []domain.Suppliers
	var Stores []domain.Store

	cursor, err := a.database.Collection(a.productCollection).Find(context.TODO(), bson.M{})
	if err != nil {
		return domain.Report{}, err
	}
	if err = cursor.All(context.TODO(), &products); err != nil {
		return domain.Report{}, err
	}

	cursor, err = a.database.Collection(a.orderCollection).Find(context.TODO(), bson.M{})
	if err != nil {
		return domain.Report{}, err
	}
	if err = cursor.All(context.TODO(), &orders); err != nil {
		return domain.Report{}, err
	}


	cursor, err = a.database.Collection(a.sellerCollection).Find(context.TODO(), bson.M{})
	if err != nil {
		return domain.Report{}, err
	}

	if err = cursor.All(context.TODO(), &Suppliers); err != nil {
		return domain.Report{}, err
	}

	cursor, err = a.database.Collection(a.storeCollection).Find(context.TODO(), bson.M{})

	if err != nil {
		return domain.Report{}, err
	}

	if err = cursor.All(context.TODO(), &Stores); err != nil {
		return domain.Report{}, err
	}

	var bestSellingProduct []domain.BestSellingProduct
	sort.Slice(products, func(i, j int) bool {
		return (products[i].BuyPrice * products[i].Quantity) > (products[j].BuyPrice * products[j].Quantity)
	})
	for i:= 0; i<5; i++ {
		bestSellingProduct = append(bestSellingProduct, domain.BestSellingProduct{
			ProductID: products[i].ID.Hex(),
			ProductName: products[i].Name,
			Category: products[i].Category,
			RemainingQuantity: products[i].Quantity,
			TurnOver: products[i].BuyPrice * products[i].Quantity,
			Price: products[i].BuyPrice,
		})

	}

	reports.BestSellingProduct = bestSellingProduct
	reports.Revenue = 0
	reports.Profit = 0
	for _, product := range products {
		reports.Revenue += product.SellPrice * product.Quantity
		reports.Profit += (product.SellPrice - product.BuyPrice) * product.Quantity
	}

	return reports, nil
}


func (a *ReportRepository) GetBestSellingCategory() ([]domain.BestSellingCategory, error) {
	var products []domain.Product
	var bestSellingCategory []domain.BestSellingCategory
	cursor, err := a.database.Collection(a.productCollection).Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &products); err != nil {
		return nil, err
	}

	categoryMap := make(map[string]int)
	productIDMap := make(map[string]string)
	for _, product := range products {
		categoryMap[product.Category] += product.SellPrice * product.Quantity
		productIDMap[product.Category] = product.ID.Hex()
	}

	for key, value := range categoryMap {
		bestSellingCategory = append(bestSellingCategory, domain.BestSellingCategory{
			Category: key,
			TurnOver: value,
			ProductID: productIDMap[key],
		})
	}

	sort.Slice(bestSellingCategory, func(i, j int) bool {
		return bestSellingCategory[i].TurnOver > bestSellingCategory[j].TurnOver
	})

	return bestSellingCategory, nil
}

func (a *ReportRepository) GetOverView() (domain.OverView, error) {
	var products []domain.Product
	var orders []domain.Orders
	var suppliers []domain.Suppliers
	var stores []domain.Store

	cursor, err := a.database.Collection(a.productCollection).Find(context.TODO(), bson.M{})
	if err != nil {
		return domain.OverView{}, err
	}
	if err = cursor.All(context.TODO(), &products); err != nil {
		return domain.OverView{}, err
	}

	cursor, err = a.database.Collection(a.orderCollection).Find(context.TODO(), bson.M{})
	if err != nil {
		return domain.OverView{}, err
	}
	if err = cursor.All(context.Background(), &orders); err != nil {
		return domain.OverView{}, err
	}

	cursor, err = a.database.Collection(a.sellerCollection).Find(context.TODO(), bson.M{})
	if err != nil {
		return domain.OverView{}, err
	}
	if err = cursor.All(context.TODO(), &suppliers); err != nil {
		return domain.OverView{}, err
	}

	cursor, err = a.database.Collection(a.storeCollection).Find(context.TODO(), bson.M{})
	if err != nil {
		return domain.OverView{}, err
	}
	if err = cursor.All(context.TODO(), &stores); err != nil {
		return domain.OverView{}, err
	}

	totalProduct := len(products)
	totalRevenue := 0
	totalProfit := 0
	totalCost := 0
	totalSold := 0
	totalCategory := 0
	totalSupplier := len(suppliers)
	topSelling := 0
	totalOrders := len(orders)

	for _, product := range products {
		totalRevenue += product.SellPrice * product.Quantity
		totalProfit += (product.SellPrice - product.BuyPrice) * product.Quantity
		totalCost += product.BuyPrice * product.Quantity
		totalSold += product.HowManySold
	}

	// for _, order := range orders {
	// 	topSelling += order.TotalPrice
	// }

	for _, product := range products {
		totalCategory += len(product.Category)
	}

	return domain.OverView{
		TotalProduct:  totalProduct,
		TotalRevenue:  totalRevenue,
		TotalProfit:   totalProfit,
		TotalCost:     totalCost,
		TotalSold:     totalSold,
		TotalCategory: totalCategory,
		TotalSupplier: totalSupplier,
		TopSelling:    topSelling,
		TotalOrders:   totalOrders,
	}, nil
}