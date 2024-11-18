package models

import "gorm.io/gorm"

type ChatHistory struct {
	gorm.Model
	SenderWspNumber string `gorm:"type:varchar(100);not null"`
	BotWspNumber    string `gorm:"type:varchar(100);not null"`
	Message         string `gorm:"type:varchar(255);not null"`
	BotResponse     string `gorm:"type:varchar(255);not null"`
	RestaurantID    uint   `gorm:"not null"`
}
