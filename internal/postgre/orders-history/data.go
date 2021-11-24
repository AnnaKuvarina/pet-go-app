package postgre

import (
	"github.com/google/uuid"
	"time"
)

type OrdersHistoryRecord struct {
	Id     uuid.UUID `gorm:"privateKey;column:id;type:uuid"`
	UserId uuid.UUID `gorm:"foreignKey;column:userId;type:uuid"`
	Date   time.Time `gorm:"column:date;type:timestamp"`
	Amount float32   `gorm:"column:amount;type:float(4)"`
	Status string    `gorm:"column:status;type:varchar(10)"`
	Item   uuid.UUID `gorm:"column:itemId;type:uuid"`
}
