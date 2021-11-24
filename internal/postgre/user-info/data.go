package postgre

import (
	"github.com/AnnaKuvarina/pet-go-app/pkg/stores"
	"github.com/google/uuid"
)

type IUserInfoStore interface {
	CreateUser(userInfo *UserInfoRecordItem) error
	UpdateUser(userInfo *UserInfoRecordItem) (*UserInfoRecordItem, error)
	DeleteUser(userId uuid.UUID) error
}

type UserInfoStore struct {
	Store *stores.PGStore
	IUserInfoStore
}

type UserInfoRecordItem struct {
	Id       uuid.UUID `gorm:"primaryKey;column:id;type:uuid"`
	Name     string    `gorm:"column:name;type:varchar(255)"`
	LastName string    `gorm:"column:lastName;type:varchar(255)"`
	Region   string    `gorm:"column:region;type:varchar(255)"`
	City     string    `gorm:"column:city;type:varchar(255)"`
	Phone    string    `gorm:"column:phone;type:varchar(13)"`
}
