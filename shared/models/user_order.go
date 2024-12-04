package models

import "gorm.io/gorm"

type UserOrder struct {
	gorm.Model
	DeliveryAddress string          `gorm:"type:varchar(255);not null"`
	UserName        string          `gorm:"type:varchar(100);not null"`
	PhoneNumber     string          `gorm:"type:varchar(20);not null"`
	PaymentMethod   string          `gorm:"type:varchar(20);not null"`
	TotalPrice      int             `gorm:"not null"`
	MenuItems       []OrderMenuItem `gorm:"foreignKey:OrderID"`
}
