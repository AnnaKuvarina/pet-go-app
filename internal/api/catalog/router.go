package catalog

import (
	"github.com/gorilla/mux"
)

func NewHttpRouter(appRouter *mux.Router, handler *Handler) {
	appRouter.HandleFunc("/products/category", handler.GetCatalogByCategory).Methods("GET")
	appRouter.HandleFunc("/products/id", handler.GetProducDetails).Methods("GET")
}
