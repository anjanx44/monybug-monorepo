package repositories

import (
	"monybug-backend/internal/database"
	"monybug-backend/internal/models"
)

func GetAllInsights() ([]models.ExpenseInsight, error) {
	var insights []models.ExpenseInsight
	result := database.GormDB.Order("generated_at DESC").Find(&insights)
	return insights, result.Error
}

func CreateInsight(i *models.ExpenseInsight) error {
	return database.GormDB.Create(i).Error
}
