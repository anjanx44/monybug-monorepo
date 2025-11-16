package repositories

import (
	"monybug-backend/internal/database"
	"monybug-backend/internal/models"
)

func GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := database.GormDB.Preload("Category").Order("transaction_date DESC").Find(&transactions)
	return transactions, result.Error
}

func CreateTransaction(t *models.Transaction) error {
	return database.GormDB.Create(t).Error
}
