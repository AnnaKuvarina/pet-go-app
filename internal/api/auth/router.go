package auth

import (
	"github.com/AnnaKuvarina/pet-go-app/pkg/api"
	"github.com/gorilla/mux"
)

func NewHttpRouter(appRouter *mux.Router, handler *Handler) {
	appRouter.HandleFunc("/login", handler.Login).Methods("POST")
	appRouter.HandleFunc("/signup", handler.Signup).Methods("POST")

	appRouter.Use(api.ContentTypeMiddleware)
}
