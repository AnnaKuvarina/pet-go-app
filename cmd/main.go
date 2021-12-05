package main

import (
	"context"
	"fmt"
	"github.com/AnnaKuvarina/pet-go-app/internal/api/auth"
	"github.com/AnnaKuvarina/pet-go-app/internal/api/catalog"
	ordershistory "github.com/AnnaKuvarina/pet-go-app/internal/api/orders-history"
	userInfo "github.com/AnnaKuvarina/pet-go-app/internal/api/user-info"
	"github.com/AnnaKuvarina/pet-go-app/internal/mongo"
	. "github.com/AnnaKuvarina/pet-go-app/internal/pg-store"
	. "github.com/AnnaKuvarina/pet-go-app/internal/services/auth"
	"github.com/AnnaKuvarina/pet-go-app/pkg/api"
	"github.com/AnnaKuvarina/pet-go-app/pkg/config"
	"github.com/AnnaKuvarina/pet-go-app/pkg/stores"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Stores struct {
	AuthStore          *UserCredsStore
	UserInfoStore      *UserInfoStore
	OrdersHistoryStore *OrdersHistoryStore
	CatalogStore       *mongo.AppMongoStore
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
	userInfoStore := NewUserInfoStore(postgreStore)
	ordersHistoryStore := NewOrdersHistoryStore(postgreStore)
	appMongoStore := mongo.NewAppMongoStore(mongoStore, config.AppConfig.MongoStore)

	return &Stores{
		CatalogStore:       appMongoStore,
		AuthStore:          authStore,
		UserInfoStore:      userInfoStore,
		OrdersHistoryStore: ordersHistoryStore,
	}, nil
}

func initRouters(appStores *Stores) *mux.Router {
	router := *mux.NewRouter()
	// router.PathPrefix("/api")
	router.Use(api.ContentTypeMiddleware)
	router.Use(api.TrimSuffixMiddleware)

	authService := NewAuthService()
	authHandler := &auth.Handler{
		Store:       appStores.AuthStore,
		AuthService: authService,
	}
	catalogHandler := &catalog.Handler{
		CatalogStore: appStores.CatalogStore,
	}
	userHandler := &userInfo.Handler{
		Store: appStores.UserInfoStore,
	}
	ordersHistoryHelper := &ordershistory.Handler{
		HistoryStore: appStores.OrdersHistoryStore,
		CatalogStore: appStores.CatalogStore,
	}
	auth.NewHttpRouter(&router, authHandler)
	catalog.NewHttpRouter(&router, catalogHandler)
	userInfo.NewHttpRouter(&router, userHandler, authHandler)
	ordershistory.NewHttpRouter(&router, ordersHistoryHelper, authHandler)

	return &router
}
