package models

type ProductCategory struct {
	ID         int `json:"id" gorm:"primary_key:auto_increment"`
	ProductID  int `json:"-"  gorm:"foreignkey:ProductID"`
	CategoryID int `json:"-"  gorm:"foreignkey:CategoryID"`
}
