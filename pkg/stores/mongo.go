package stores

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	DBUrl              string
	DBName             string
	ProductsCollection string
	CommentsCollection string
}

type MongoStore struct {
	Client *mongo.Client
}

func NewMongoDBStore(dbConfig *MongoConfig) (*MongoStore, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbConfig.DBUrl))
	if err != nil {
		log.Fatal().Msg(err.Error())
		return nil, err
	}

	store := &MongoStore{
		Client: client,
	}

	return store, nil
}

func (s *MongoStore) Close(ctx context.Context) {
	err := s.Client.Disconnect(ctx)

	if err != nil {
		log.Err(err).Msgf("Failed to close connection to MongoDB")
	}

	fmt.Println("Connection to MongoDB closed.")
}

func (s *MongoStore) Connect(ctx context.Context) error {
	// Create connect
	err := s.Client.Connect(context.TODO())
	if err != nil {
		fmt.Println("Failed to connect to MongoDB")
		return err
	}

	// Check the connection
	err = s.Client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Failed to check connection to MongoDB")
		return err
	}

	return nil
}
