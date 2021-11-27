package postgre

import (
	"context"
	"github.com/AnnaKuvarina/pet-go-app/pkg/stores"
)

func NewUserCredsStore(pgStore *stores.PGStore) *UserCredsStore {
	return &UserCredsStore{
		Store: pgStore,
	}
}

func GetUserByEmail(ctx context.Context,email string) (*UserCredItem, error){
	return nil, nil
}
