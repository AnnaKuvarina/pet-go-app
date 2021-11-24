package mongo

import (
	"github.com/AnnaKuvarina/pet-go-app/pkg/stores"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppMongoStore struct {
	Store *stores.MongoStore
	collections []*mongo.Collection
}

func NewAppMongoStore(store *stores.MongoStore, cfg *stores.MongoConfig) *AppMongoStore {
	products := store.Client.Database(cfg.DBName).Collection(cfg.ProductsCollection)
	comments := store.Client.Database(cfg.DBName).Collection(cfg.CommentsCollection)

	var collections []*mongo.Collection
	collections = append(collections, products)
	collections = append(collections, comments)

	return &AppMongoStore{
		Store:      store,
		collections: collections,
	}
}
