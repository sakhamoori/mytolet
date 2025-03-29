package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/sakhamoori/mytolet/api/internal/models"
)

func InitDB() (*gorm.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/mytolet_db"
	}

	config := &gorm.Config{}
	
	if os.Getenv("ENV") == "development" {
		config.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(dbURL), config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto Migrate the models
	err = db.AutoMigrate(
		&models.User{},
		&models.Property{},
		&models.Address{},
		&models.Application{},
		&models.Document{},
		&models.Message{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to auto-migrate database: %w", err)
	}

	log.Println("Database connection established")
	return db, nil
}