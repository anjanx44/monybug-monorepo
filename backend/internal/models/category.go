package models

type Category struct {
	ID       uint   `json:"id" gorm:"primaryKey;column:category_id"`
	Name     string `json:"name" gorm:"uniqueIndex;not null"`
	Type     string `json:"type" gorm:"check:type IN ('INCOME','EXPENSE')"`
	Priority string `json:"priority"`
	Color    string `json:"color_code" gorm:"column:color_code"`
	IsActive bool   `json:"is_active" gorm:"default:true"`
}
