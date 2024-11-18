package models

import "gorm.io/gorm"

type ChatHistory struct {
	gorm.Model
	SenderWspNumber string     `gorm:"type:varchar(20);not null"`
	BotWspNumber    string     `gorm:"type:varchar(20);not null"`
	Message         string     `gorm:"type:text;not null"`
	BotResponse     string     `gorm:"type:text;not null"`
	RestaurantID    uint       `gorm:"not null;index"`
	Restaurant      Restaurant `gorm:"foreignKey:RestaurantID"`
}
