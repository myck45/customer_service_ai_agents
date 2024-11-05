package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);unique;not null"`
	Menu []Menu `gorm:"foreignKey:RestaurantID"`
}
