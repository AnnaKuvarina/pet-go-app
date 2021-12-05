package pg_store

import "github.com/AnnaKuvarina/pet-go-app/pkg/stores"

type IOrdersHistoryStore interface {
}

type OrdersHistoryStore struct {
	Store *stores.PGStore
	IOrdersHistoryStore
}

func NewOrdersHistoryStore(pgStore *stores.PGStore) *OrdersHistoryStore {
	return &OrdersHistoryStore{
		Store: pgStore,
	}
}
