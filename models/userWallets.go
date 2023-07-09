// models/user_wallet.go
package models

import (
	"gorm.io/gorm"
)

type UserWallet struct {
	gorm.Model
	UserID        string `gorm:"unique;not null"`
	WalletAddress string `gorm:"unique;not null"`
}

func CreateUserWallet(userWallet *UserWallet) error {
	result := DB.Create(userWallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserWalletByID(id uint) (*UserWallet, error) {
	var userWallet UserWallet
	result := DB.First(&userWallet, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userWallet, nil
}

func GetUserWalletByUserID(userID string) (*UserWallet, error) {
	var userWallet UserWallet
	result := DB.Where("user_id = ?", userID).First(&userWallet)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userWallet, nil
}

func GetAllUserWallets() ([]UserWallet, error) {
	var userWallets []UserWallet
	result := DB.Find(&userWallets)
	if result.Error != nil {
		return nil, result.Error
	}
	return userWallets, nil
}

func UpdateUserWallet(userWallet *UserWallet) error {
	result := DB.Save(userWallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteUserWallet(userWallet *UserWallet) error {
	result := DB.Delete(userWallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
