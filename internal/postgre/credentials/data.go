package postgre

import (
	"github.com/AnnaKuvarina/pet-go-app/pkg/stores"
)

type IUserCredsStore interface {
	Create(user *UserCredItem) error
	GetPasswordByUserName(userName string) (string, error)
	FindUserByEmail(email string) (*UserCredItem, error)
	UpdatePassword()
}

type UserCredsStore struct {
	Store *stores.PGStore
	IUserCredsStore
}

type UserCredItem struct {
	stores.TableItem
	ID        string `gorm:"primaryKey;column:userId;type:uuid"`
	Username  string `gorm:"column:username;type:varchar(255)"`
	Password  string `gorm:"column:password;type:text"`
	Email     string `gorm:"unique;column:email;type:varchar(255)"`
	TokenHash string `gorm:"not null;column:tokenHash;type:varchar(15)"`
}
