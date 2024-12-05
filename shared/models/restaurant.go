package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	Name          string        `gorm:"type:varchar(100);unique;not null"`
	Menus         []Menu        `gorm:"foreignKey:RestaurantID"`
	ChatHistories []ChatHistory `gorm:"foreignKey:RestaurantID"`
	Bot           []Bot         `gorm:"foreignKey:RestaurantID"`
	UserOrder     []UserOrder   `gorm:"foreignKey:RestaurantID"`
	UserID        uint          `gorm:"not null;index"`
	User          User          `gorm:"foreignKey:UserID"`
}
