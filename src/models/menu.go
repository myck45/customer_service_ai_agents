package models

import (
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	RestaurantID uint            `gorm:"not null"`
	ItemName     string          `gorm:"type:varchar(100);not null"`
	Description  string          `gorm:"type:varchar(255);not null"`
	Price        int             `gorm:"not null"`
	Embedding    pgvector.Vector `gorm:"type:vector(3072);index:idx_embedding,type:ivfflat"`
}
