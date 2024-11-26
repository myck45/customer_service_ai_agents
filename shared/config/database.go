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

	gormConfig := &gorm.Config{
		PrepareStmt: false,
	}

	db, err := gorm.Open(postgres.Open(sqlInfo), gormConfig)
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
		logrus.Warnf("Error migrating models: %v", err)
	}

	// Load SQL functions
	searchMenuSQL := SearchMenuSQLFunction()
	err = db.Exec(searchMenuSQL).Error
	if err != nil {
		logrus.Warnf("Error executing SQL file: %v", err)
	}

	return db
}

func SearchMenuSQLFunction() string {
	return `
		CREATE OR REPLACE FUNCTION search_menu(
    query_embedding vector(3072),
    similarity_threshold float,
    match_count int,
    restaurant_id bigint
)
RETURNS TABLE (
    id bigint,
    item_name text,
    price int,
    description text,
    likes int,
    similarity float
)
LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY
    SELECT
        menus.id,
        menus.item_name,
        menus.price,
        menus.description,
        menus.likes,
        menus.embedding <#> query_embedding AS similarity
    FROM
        menus
    WHERE
        menus.restaurant_id = restaurant_id
        AND menus.embedding <#> query_embedding < similarity_threshold
    ORDER BY
        similarity
    LIMIT
        match_count;
END;
$$;
	`
}
