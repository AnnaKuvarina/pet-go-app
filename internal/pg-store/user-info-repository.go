package pg_store

import "github.com/AnnaKuvarina/pet-go-app/pkg/stores"

func NewUserInfoStore(pgStore *stores.PGStore) *UserInfoStore {
	return &UserInfoStore{
		Store: pgStore,
	}
}
