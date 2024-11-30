package repository

import (
	"context"
	"inventory/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StoreRepository struct {
	database   mongo.Database
	collection string
}

func NewStoreRepository(database mongo.Database, collection string) domain.StoreRepository {
	return &StoreRepository{
		database:   database,
		collection: collection}

}

func (a *StoreRepository) CreateStore(store domain.Store) (domain.Store, error) {
	objID := primitive.NewObjectID()
	store.ID = objID
	_, err := a.database.Collection(a.collection).InsertOne(context.Background(), store)
	if err != nil {
		return domain.Store{}, err
	}
	return store, nil
}

func (a *StoreRepository) GetAllStore() ([]domain.Store, error) {
	var store []domain.Store
	cursor, err := a.database.Collection(a.collection).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &store); err != nil {
		return nil, err
	}
	return store, nil
}


func (a *StoreRepository) UpdateStore(store domain.Store) (domain.Store, error) {
	filter := bson.M{"_id": store.ID}
	update := bson.M{"$set": store}
	_, err := a.database.Collection(a.collection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return domain.Store{}, err
	}
	return store, nil
}

func (a *StoreRepository) DeleteStore(id string) (domain.Store, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Store{}, err
	}
	filter := bson.M{"_id": objID}
	_, err = a.database.Collection(a.collection).DeleteOne(context.Background(), filter)
	if err != nil {
		return domain.Store{}, err
	}
	return domain.Store{}, nil
}

func (a *StoreRepository) GetStoreByID(id string) (domain.Store, error) {
	var store domain.Store
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Store{}, err
	}
	filter := bson.M{"_id": objID}
	err = a.database.Collection(a.collection).FindOne(context.Background(), filter).Decode(&store)
	if err != nil {
		return domain.Store{}, err
	}
	return store, nil
}