package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Contract struct {
	gorm.Model
	Network      string `gorm:"not null"`
	Address      string `gorm:"unique;not null"`
	ABI          string `gorm:"type:text;not null"`
}

func (c *Contract) CreateContract() error {
	if DB != nil {
		fmt.Println("DB is null")
	}
	result := DB.Create(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetContractByID(id uint) (*Contract, error) {
	var contract Contract
	result := DB.First(&contract, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &contract, nil
}

func GetContractByAddress(address string) (*Contract, error) {
	var contract Contract
	result := DB.Where("address = ?", address).First(&contract)
	if result.Error != nil {
		return nil, result.Error
	}
	return &contract, nil
}

func GetAllContracts() ([]Contract, error) {
	var contracts []Contract
	result := DB.Find(&contracts)
	if result.Error != nil {
		return nil, result.Error
	}
	return contracts, nil
}

func (c *Contract) UpdateContract() error {
	result := DB.Save(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *Contract) DeleteContract() error {
	result := DB.Delete(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
