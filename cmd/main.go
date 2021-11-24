package main

import (
	"context"
	"fmt"
	"github.com/AnnaKuvarina/pet-go-app/internal/mongo"
	"github.com/AnnaKuvarina/pet-go-app/internal/postgre"
	"github.com/AnnaKuvarina/pet-go-app/internal/products"
	"github.com/AnnaKuvarina/pet-go-app/internal/users"
	"github.com/AnnaKuvarina/pet-go-app/pkg/config"
	"github.com/AnnaKuvarina/pet-go-app/pkg/stores"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

var (
	appPostgreStore *postgre.AppPostgreStore
	appMongoStore   *mongo.AppMongoStore
	productsRouter  *mux.Router
	usersRouter     *mux.Router
)

func main() {
	log.Debug().Msg("Start server")
	err := LoadConfig()
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("can't load config")
		return
	}


}

func InitApp() error {
	ctx := context.Background()

	postgreStore, err := stores.NewPostgreStore(config.AppConfig.PostgreStore)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create PostgreSQL store")
	}

	mongoStore, err := stores.NewMongoDBStore(config.AppConfig.MongoStore)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create MongoDB store")
	}

	err = mongoStore.Connect(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to MongoDB store")
	}
	fmt.Println("connected to MongoDB")

	appPostgreStore = postgre.NewAppStore(postgreStore)
	appMongoStore = mongo.NewAppMongoStore(mongoStore, config.AppConfig.MongoStore)

	usersRouter = users.NewUsersRouter(appMongoStore, appPostgreStore)
	productsRouter = products.NewProductsRouter(appMongoStore)

	return nil
}
