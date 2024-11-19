package models

import (
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	ItemName     string          `gorm:"type:varchar(100);not null"`
	Description  string          `gorm:"type:varchar(255);not null"`
	Price        int             `gorm:"not null"`
	Likes        int             `gorm:"default:0"`
	Embedding    pgvector.Vector `gorm:"type:vector(3072)"`
	RestaurantID uint            `gorm:"not null;index"`
	Restaurant   Restaurant      `gorm:"foreignKey:RestaurantID"`
}
