package models

import (
	"gorm.io/gorm"
)
type Transaction struct {
	gorm.Model
	HashorResult       string `gorm:"column:hash_or_result"`
	UserID             string `gorm:"column:user_id"`
	Network            string `gorm:"column:network"`
	TargetContract     string `gorm:"column:target_contract"`
	TargetFunction     string `gorm:"column:target_function"`
	Args               string `gorm:"column:args"`
	IsMetaTransaction  bool `gorm:"column:is_meta_transaction"`
	TransactionType    MethodType `gorm:"column:transaction_type"`
	IsSuccessful       bool `gorm:"column:is_successful"`
	IsTestnet          bool `gorm:"column:is_testnet"`
}

type MethodType string

const (
	MethodTypeRead  MethodType = "read"
	MethodTypeWrite MethodType = "write"
)

func CreateTransaction(transaction *Transaction) error {
	result := DB.Create(transaction)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetTransactionByID(id uint) (*Transaction, error) {
	var transaction Transaction
	result := DB.First(&transaction, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &transaction, nil
}

func GetAllTransactions() ([]Transaction, error) {
	var transactions []Transaction
	result := DB.Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func UpdateTransaction(transaction *Transaction) error {
	result := DB.Save(transaction)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteTransaction(transaction *Transaction) error {
	result := DB.Delete(transaction)
	if result.Error != nil {
		return result.Error
	}
	return nil
}