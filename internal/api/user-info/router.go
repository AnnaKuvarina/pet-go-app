package user_info

import (
	"github.com/AnnaKuvarina/pet-go-app/internal/api/auth"
	"github.com/gorilla/mux"
)

func NewHttpRouter(appRouter *mux.Router, handler *Handler, authHandler *auth.Handler) {
	userInfoRouter := appRouter.PathPrefix("/users").Subrouter()
	userInfoRouter.HandleFunc("/:id", handler.GetUserData).Methods("GET")
	userInfoRouter.Use(authHandler.MiddlewareValidateAccessToken)
}