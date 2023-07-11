package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
    Username  string         `gorm:"unique;not null"`
    Email     string         `gorm:"unique;not null"`
    Password  string         `gorm:"not null"`
    Role      string         `gorm:"not null"`
    CreatedAt time.Time      `gorm:"autoCreateTime"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime"`
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateUser(user *User) error {
	result := DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserByID(id uint) (*User, error) {
	var user User
	result := DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	result := DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	result := DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func UpdateUser(user *User) error {
	result := DB.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteUser(user *User) error {
	result := DB.Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUserByUsernameOrEmail retrieves a user by username or email
func GetUserByUsernameOrEmail(username string) (*User, error) {
	var user User
	result := DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
