package main

import (
	"context"
	"fmt"
	"github.com/AnnaKuvarina/pet-go-app/internal/api/auth"
	"github.com/AnnaKuvarina/pet-go-app/internal/mongo"
	. "github.com/AnnaKuvarina/pet-go-app/internal/postgre/credentials"
	. "github.com/AnnaKuvarina/pet-go-app/internal/services/auth"
	"github.com/AnnaKuvarina/pet-go-app/pkg/config"
	"github.com/AnnaKuvarina/pet-go-app/pkg/stores"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Stores struct {
	AuthStore    *UserCredsStore
	CatalogStore *mongo.AppMongoStore
}

func main() {
	log.Debug().Msg("Start server")
	err := LoadConfig()
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("can't load config")
		return
	}

	appStores, err := InitStores()
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("failed to init stores")
		return
	}

	router := initRouters(appStores)

	errChan := make(chan error)

	go func() {
		errChan <- http.ListenAndServe(fmt.Sprintf(":%d", config.AppConfig.Port), router)
	}()

}

func InitStores() (*Stores, error) {
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

	authStore := NewUserCredsStore(postgreStore)
	appMongoStore := mongo.NewAppMongoStore(mongoStore, config.AppConfig.MongoStore)

	return &Stores{
		CatalogStore: appMongoStore,
		AuthStore: authStore,
	}, nil
}

func initRouters(appStores *Stores) *mux.Router {
	router := *mux.NewRouter()
	authService := NewAuthService()
	authHandler := &auth.Handler{
		Store:       appStores.AuthStore,
		AuthService: authService,
	}
	auth.NewHttpRouter(&router, authHandler)

	return &router
}
