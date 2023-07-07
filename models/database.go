package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db

	// Auto-migrate the contract model
	err = db.AutoMigrate(&Contract{})
	if err != nil {
		return fmt.Errorf("failed to perform auto migration: %w", err)
	}

	return nil
}
