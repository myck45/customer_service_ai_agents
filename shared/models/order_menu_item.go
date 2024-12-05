package models

import "gorm.io/gorm"

type OrderMenuItem struct {
	gorm.Model
	ItemName    string    `gorm:"type:varchar(100);not null"`
	Quantity    int       `gorm:"not null"`
	Price       int       `gorm:"not null"`
	Subtotal    int       `gorm:"not null"`
	UserOrderID uint      `gorm:"not null;index"`
	UserOrder   UserOrder `gorm:"foreignKey:UserOrderID"`
}
