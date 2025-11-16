package models

import "time"

type Budget struct {
	ID           int       `json:"id" db:"budget_id"`
	CategoryID   int       `json:"category_id" db:"category_id"`
	MonthlyLimit float64   `json:"monthly_limit" db:"monthly_limit"`
	Month        time.Time `json:"month" db:"month"`
	IsActive     bool      `json:"is_active" db:"is_active"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
