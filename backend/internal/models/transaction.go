package models

import "time"

type Transaction struct {
	ID          int       `json:"id" db:"transaction_id"`
	CategoryID  int       `json:"category_id" db:"category_id"`
	Date        time.Time `json:"date" db:"transaction_date"`
	Amount      float64   `json:"amount" db:"amount"`
	Currency    string    `json:"currency" db:"currency"`
	Description string    `json:"description" db:"description"`
	Tags        string    `json:"tags" db:"tags"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
