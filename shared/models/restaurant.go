package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	UserID      uint          `gorm:"not null"`
	Name        string        `gorm:"type:varchar(100);unique;not null"`
	Menu        []Menu        `gorm:"foreignKey:RestaurantID"`
	ChatHistory []ChatHistory `gorm:"foreignKey:RestaurantID"`
	Bot         []Bot         `gorm:"foreignKey:RestaurantID"`
}
