package models

type ProductCategory struct {
	ID         int `json:"id" gorm:"primary_key:auto_increment"`
	ProductID  int `json:"product_id"`
	CategoryID int `json:"category_id"`
}
