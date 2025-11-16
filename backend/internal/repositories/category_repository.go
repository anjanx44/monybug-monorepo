package repositories

import (
	"monybug-backend/internal/database"
	"monybug-backend/internal/models"
)

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	result := database.GormDB.Where("is_active = ?", true).Find(&categories)
	return categories, result.Error
}

func CreateCategory(c *models.Category) error {
	return database.GormDB.Create(c).Error
}
