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

func UpdateCategory(id uint, c *models.Category) error {
	return database.GormDB.Where("category_id = ?", id).Updates(c).Error
}

func DeleteCategory(id uint) error {
	return database.GormDB.Model(&models.Category{}).Where("category_id = ?", id).Update("is_active", false).Error
}
