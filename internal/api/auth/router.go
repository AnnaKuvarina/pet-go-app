package auth

import "github.com/gorilla/mux"

func NewHttpRouter(handler *Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/signup", handler.Signup).Methods("POST")

	return router
}
