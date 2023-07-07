package utils

import (
	"fmt"

	"github.com/Jahankohan/mpc_wallet/config"
	"github.com/Jahankohan/mpc_wallet/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(dbConfig config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.DBName, dbConfig.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	fmt.Println("DB:", db)

	err = db.AutoMigrate(&models.Contract{})
	if err != nil {
		return nil, fmt.Errorf("failed to perform auto migration: %w", err)
	}

	return db, nil
}
