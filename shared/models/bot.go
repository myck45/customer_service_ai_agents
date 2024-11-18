package models

import "gorm.io/gorm"

type Bot struct {
	gorm.Model
	Name         string     `gorm:"type:varchar(100);unique;not null"`
	Identity     string     `gorm:"type:varchar(100);not null"`
	WspNumber    string     `gorm:"type:varchar(20);unique;not null"`
	RestaurantID uint       `gorm:"not null;index"`
	Restaurant   Restaurant `gorm:"foreignKey:RestaurantID"`
}
