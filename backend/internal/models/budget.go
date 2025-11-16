package models

import "time"

type Budget struct {
	ID           uint      `json:"id" gorm:"primaryKey;column:budget_id"`
	CategoryID   uint      `json:"category_id" gorm:"not null"`
	Category     Category  `json:"category" gorm:"foreignKey:CategoryID"`
	MonthlyLimit float64   `json:"monthly_limit" gorm:"type:decimal(10,2);not null"`
	Month        time.Time `json:"month" gorm:"not null"`
	IsActive     bool      `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time `json:"created_at"`
}
