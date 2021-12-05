package ordershistory

import (
	"github.com/AnnaKuvarina/pet-go-app/internal/api/auth"
	"github.com/gorilla/mux"
)

func NewHttpRouter(appRouter *mux.Router, handler *Handler, authHandler *auth.Handler) {
	historyRouter := appRouter.PathPrefix("/orders").Subrouter()
	historyRouter.HandleFunc("/history/user/:id", handler.GetUserOrdersHistory).Methods("GET")
	historyRouter.HandleFunc("/new", handler.NewOrder).Methods("POST")
	historyRouter.Use(authHandler.MiddlewareValidateAccessToken)
}