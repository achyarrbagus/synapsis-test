package models

import "time"

type Product struct {
	ID          int        `json:"id" gorm:"primary_key:auto_increment"`
	Name        string     `json:"name" form:"name" gorm:"type: varchar(255)"`
	Description string     `json:"description" form:"description" gorm:"type: varchar(255)"`
	Price       int        `json:"price" form:"price" gorm:"type: int"`
	Stock       int        `json:"stock" form:"stock" gorm:"type: int"`
	Image       string     `json:"image" form:"image" gorm:"type: varchar(255)"`
	Category    []Category `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;many2many:product_categories"`
	CategoryID  []int      `json:"-" form:"category_id" gorm:"-"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
