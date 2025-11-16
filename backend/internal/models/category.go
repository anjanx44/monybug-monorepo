package models

type Category struct {
	ID       int    `json:"id" db:"category_id"`
	Name     string `json:"name" db:"name"`
	Type     string `json:"type" db:"type"`
	Priority string `json:"priority" db:"priority"`
	Color    string `json:"color_code" db:"color_code"`
	IsActive bool   `json:"is_active" db:"is_active"`
}
