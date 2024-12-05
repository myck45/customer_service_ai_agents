package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string       `gorm:"type:varchar(100);not null"`
	LastName    string       `gorm:"type:varchar(100);not null"`
	BirthDate   time.Time    `gorm:"type:date;not null"`
	UserEmail   string       `gorm:"type:varchar(100);unique;not null"`
	Password    string       `gorm:"type:varchar(255);not null"`
	PhoneNum    string       `gorm:"type:varchar(25);unique;not null"`
	Role        string       `gorm:"type:varchar(25);not null"`
	Restaurants []Restaurant `gorm:"foreignKey:UserID"`
}
