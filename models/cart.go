package models

import "time"

type Cart struct {
	ID            int         `json:"id" gorm:"primary_key:auto_increment"`
	UserID        int         `json:"user_id"`
	User          User        `json:"user" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;foreignKey:UserID"`
	ProductID     int         `json:"product_id"`
	Product       Product     `json:"cart" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;foreignKey:ProductID"`
	Qty           int         `json:"qty" gorm:"type:int"`
	TransactionID int         `json:"transaction_id" gorm:"type: int"`
	Transaction   Transaction `json:"transaction"`
	CreatedAt     time.Time   `json:"-"`
	UpdatedAt     time.Time   `json:"-"`
}
