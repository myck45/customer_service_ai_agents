package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string       `gorm:"type:varchar(100);unique;not null"`
	UserEmail  string       `gorm:"type:varchar(100);unique;not null"`
	Password   string       `gorm:"type:varchar(100);not null"`
	PhoneNum   string       `gorm:"type:varchar(100);unique;not null"`
	Role       string       `gorm:"type:varchar(100);not null"`
	Restaurant []Restaurant `gorm:"foreignKey:UserID"`
}
