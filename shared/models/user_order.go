package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserOrder struct {
	gorm.Model
	OrderCode       uuid.UUID       `gorm:"type:uuid;not null;unique"`
	DeliveryAddress string          `gorm:"type:varchar(255);not null"`
	UserName        string          `gorm:"type:varchar(100);not null"`
	PhoneNumber     string          `gorm:"type:varchar(20);not null"`
	PaymentMethod   string          `gorm:"type:varchar(20);not null"`
	TotalPrice      int             `gorm:"not null"`
	BotWspNumber    string          `gorm:"type:varchar(20);not null"`
	SenderWspNumber string          `gorm:"type:varchar(20);not null"`
	Status          string          `gorm:"type:varchar(20);not null"`
	RestaurantID    uint            `gorm:"not null;index"`
	Restaurant      Restaurant      `gorm:"foreignKey:RestaurantID"`
	OrderMenuItems  []OrderMenuItem `gorm:"foreignKey:UserOrderID"`
}
