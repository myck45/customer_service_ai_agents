package models

import "gorm.io/gorm"

type MenuFile struct {
	gorm.Model
	FileName     string     `gorm:"type:varchar(255);not null"`
	FilePath     string     `gorm:"type:varchar(500);not null"`
	FileType     string     `gorm:"type:varchar(50);not null"`
	FileSize     int64      `gorm:"not null"`
	RestaurantID uint       `gorm:"not null;index"`
	Restaurant   Restaurant `gorm:"foreignKey:RestaurantID"`
}
