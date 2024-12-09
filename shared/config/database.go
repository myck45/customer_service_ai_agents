package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/proyectos01-a/shared/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func DatabaseConnection() *gorm.DB {
	once.Do(func() {
		host := os.Getenv("DB_HOST")
		port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")

		sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
			host, port, user, password, dbname)

		gormConfig := &gorm.Config{
			PrepareStmt:          false,
			DisableAutomaticPing: true,
		}

		var err error
		db, err = gorm.Open(postgres.Open(sqlInfo), gormConfig)
		if err != nil {
			logrus.Fatalf("Error connecting to database: %v", err)
			panic("Failed to connect to database")
		}

		// Pool configuration
		sqlDB, err := db.DB()
		if err != nil {
			logrus.Fatalf("Error getting database connection: %v", err)
		}
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetConnMaxLifetime(time.Minute * 5)

		if err = db.Exec("CREATE EXTENSION IF NOT EXISTS vector").Error; err != nil {
			if strings.Contains(err.Error(), "already exists") {
				logrus.Warnf("Extension vector already exists")
			} else {
				logrus.Fatalf("Error creating extension vector: %v", err)
			}
		}

		if err = db.AutoMigrate(
			&models.User{},
			&models.Restaurant{},
			&models.Bot{},
			&models.Menu{},
			&models.ChatHistory{},
			&models.MenuFile{},
			&models.UserOrder{},
			&models.OrderMenuItem{},
		); err != nil {
			logrus.Warnf("Error migrating models: %v", err)
		}
	})

	return db
}
