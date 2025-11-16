package models

import "time"

type ExpenseInsight struct {
	ID          int       `json:"id" db:"insight_id"`
	Type        string    `json:"insight_type" db:"insight_type"`
	Message     string    `json:"message" db:"message"`
	Period      string    `json:"period" db:"period"`
	Data        string    `json:"data" db:"data"`
	GeneratedAt time.Time `json:"generated_at" db:"generated_at"`
}
