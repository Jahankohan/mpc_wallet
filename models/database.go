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

	// Auto-migrate the user model
	err = db.AutoMigrate(&User{})
	if err != nil {
		return fmt.Errorf("failed to perform auto migration: %w", err)
	}

	// Auto-migrate the contract model
	err = db.AutoMigrate(&Contract{})
	if err != nil {
		return fmt.Errorf("failed to perform auto migration: %w", err)
	}

	// Auto-migrate the User wallet model
	err = db.AutoMigrate(&UserWallet{})
	if err != nil {
		return fmt.Errorf("failed to perform auto migration: %w", err)
	}

	// Auto-migrate the transaction model
	err = db.AutoMigrate(&Transaction{})
	if err != nil {
		return fmt.Errorf("failed to perform auto migration: %w", err)
	}


	return nil
}
