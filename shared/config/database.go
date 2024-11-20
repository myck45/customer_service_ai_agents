package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/proyectos01-a/shared/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("Error connecting to database: %v", err)
		panic("Failed to connect to database")
	}

	err = db.Exec("CREATE EXTENSION IF NOT EXISTS vector").Error
	if err != nil {
		logrus.Fatalf("Error enabling vector extension: %v", err)
		panic("Failed to enable vector extension")
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Restaurant{},
		&models.Bot{},
		&models.Menu{},
		&models.ChatHistory{},
	)

	if err != nil {
		logrus.Fatalf("Error migrating models: %v", err)
		panic("Failed to migrate models")
	}

	return db
}
