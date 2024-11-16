package models

import "gorm.io/gorm"

type Bot struct {
	gorm.Model
	Name         string `gorm:"type:varchar(100);unique;not null"`
	WspNumber    string `gorm:"type:varchar(100);unique;not null"`
	RestaurantID uint   `gorm:"not null"`
}
