package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Store struct {
	ID          primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Name        string             `json:"name" bson:"name"`
	Location    Location       `json:"location" bson:"location"`
	Image       string             `json:"image" bson:"image"`
}

type Location struct {
	City	string `json:"city" bson:"city"`
	Subcity string `json:"subcity" bson:"subcity"`
	Kebele string `json:"kebele" bson:"kebele"`
}

type StoreUseCase interface {
	CreateStore(store Store) (Store, error)
	GetAllStore() ([]Store, error)
	GetStoreByID(id string) (Store, error)
	UpdateStore(store Store) (Store, error)
	DeleteStore(id string) (Store, error)
}

type StoreRepository interface {
	CreateStore(store Store) (Store, error)
	GetAllStore() ([]Store, error)
	GetStoreByID(id string) (Store, error)
	UpdateStore(store Store) (Store, error)
	DeleteStore(id string) (Store, error)
}	
