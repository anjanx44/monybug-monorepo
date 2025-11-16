package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID          uint           `json:"id" gorm:"primaryKey;column:transaction_id"`
	CategoryID  uint           `json:"category_id" gorm:"not null"`
	Category    Category       `json:"category" gorm:"foreignKey:CategoryID"`
	Date        time.Time      `json:"date" gorm:"column:transaction_date;not null"`
	Amount      float64        `json:"amount" gorm:"type:decimal(10,2);not null"`
	Currency    string         `json:"currency" gorm:"default:BDT"`
	Description string         `json:"description"`
	Tags        string         `json:"tags"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
