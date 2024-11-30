package repository

import (
	"context"
	"inventory/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	database   mongo.Database
	collection string
}

func NewOrderRepository(database mongo.Database, collection string) domain.OrdersRepository {
	return &OrderRepository{
		database:   database,
		collection: collection}

}

func (a *OrderRepository) CreateOrders(orders domain.Orders) (domain.Orders, error) {
	objID := primitive.NewObjectID()
	orders.ID = objID
	_, err := a.database.Collection(a.collection).InsertOne(context.Background(), orders)
	if err != nil {
		return domain.Orders{}, err
	}
	return orders, nil
}

func (a *OrderRepository) GetAllOrders() ([]domain.Orders, error) {
	var orders []domain.Orders
	cursor, err := a.database.Collection(a.collection).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &orders); err != nil {
		return nil, err
	}
	return orders, nil
}


func (a *OrderRepository) UpdateOrders(orders domain.Orders) (domain.Orders, error) {
	filter := bson.M{"_id": orders.ID}
	update := bson.M{"$set": orders}
	_, err := a.database.Collection(a.collection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return domain.Orders{}, err
	}
	return orders, nil
}

func (a *OrderRepository) DeleteOrders(id string) (domain.Orders, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Orders{}, err
	}
	_, err = a.database.Collection(a.collection).DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return domain.Orders{}, err
	}
	return domain.Orders{}, nil
}

func (a *OrderRepository) GetOrdersByID(id string) (domain.Orders, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return  domain.Orders{} , err
	}
	var orders domain.Orders
	err = a.database.Collection(a.collection).FindOne(context.Background(), bson.M{"_id": objID}).Decode(&orders)
	if err != nil {
		return domain.Orders{}, err
	}
	return orders, nil
}



