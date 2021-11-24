package postgre

import "github.com/AnnaKuvarina/pet-go-app/pkg/stores"

type IAppPostgreStore interface {
}

type AppPostgreStore struct {
	Store *stores.PGStore
	IAppPostgreStore
}

func NewAppStore(pgStore *stores.PGStore) *AppPostgreStore {
	return &AppPostgreStore{
		Store: pgStore,
	}
}

