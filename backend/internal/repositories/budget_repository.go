package repositories

import (
	"monybug-backend/internal/database"
	"monybug-backend/internal/models"
)

func GetAllBudgets() ([]models.Budget, error) {
	var budgets []models.Budget
	result := database.GormDB.Preload("Category").Where("is_active = ?", true).Find(&budgets)
	return budgets, result.Error
}

func CreateBudget(b *models.Budget) error {
	return database.GormDB.Create(b).Error
}

func UpdateBudget(id uint, b *models.Budget) error {
	return database.GormDB.Where("budget_id = ?", id).Updates(b).Error
}

func DeleteBudget(id uint) error {
	return database.GormDB.Model(&models.Budget{}).Where("budget_id = ?", id).Update("is_active", false).Error
}
