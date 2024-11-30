package repository

import (
	"context"
	"inventory/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SupplierRepository struct {
	database   mongo.Database
	collection string
}

func NewSupplierRepository(database mongo.Database, collection string) domain.SuppliersRepository {
	return &SupplierRepository{
		database:   database,
		collection: collection}

}

func (s *SupplierRepository) CreateSuppliers(suppliers domain.Suppliers) (domain.Suppliers, error) {
	objID := primitive.NewObjectID()
	suppliers.ID = objID
	_, err := s.database.Collection(s.collection).InsertOne(context.Background(), suppliers)
	if err != nil {
		return domain.Suppliers{}, err
	}
	return suppliers, nil
}

func (s *SupplierRepository) GetAllSuppliers() ([]domain.Suppliers, error) {
	var suppliers []domain.Suppliers
	cursor, err := s.database.Collection(s.collection).Find(context.Background(),  bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &suppliers); err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (s *SupplierRepository) GetSuppliersByID(id string) (domain.Suppliers, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return  domain.Suppliers{} , err
	}

	var supplier domain.Suppliers
	err = s.database.Collection(s.collection).FindOne(context.Background(), bson.M{"_id": objID}).Decode(&supplier)
	if err != nil {
		return domain.Suppliers{}, err
	}
	return supplier, nil
}

func (s *SupplierRepository) DeleteSuppliers(id string) (domain.Suppliers, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return  domain.Suppliers{} , err
	}
	var supplier domain.Suppliers
	_, err = s.database.Collection(s.collection).DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return domain.Suppliers{}, err
	}
	return supplier, nil
}


func (s *SupplierRepository) UpdateSuppliers(suppliers domain.Suppliers) (domain.Suppliers, error) {
	filter := bson.M{"_id": suppliers.ID}
	update := bson.M{"$set": suppliers}
	_, err := s.database.Collection(s.collection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return domain.Suppliers{}, err
	}
	return suppliers, nil
}