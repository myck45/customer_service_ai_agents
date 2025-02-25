package models

import "gorm.io/gorm"

type Bot struct {
	gorm.Model
	Name         string     `gorm:"type:varchar(100);unique;not null"`
	Identity     string     `gorm:"type:text;not null"`
	WspNumber    string     `gorm:"type:varchar(25);unique;not null"`
	RestaurantID uint       `gorm:"not null;index"`
	Restaurant   Restaurant `gorm:"foreignKey:RestaurantID"`
}
