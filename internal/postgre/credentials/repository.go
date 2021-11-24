package postgre

import "github.com/AnnaKuvarina/pet-go-app/pkg/stores"

func NewUserCredsStore(pgStore *stores.PGStore) *UserCredsStore {
	return &UserCredsStore{
		Store: pgStore,
	}
}

