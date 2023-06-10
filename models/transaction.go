package models

import "time"

type Transaction struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	UserID      int       `json:"-"`
	User        User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AddressID   int       `json:"addrees_id"`
	Address     Addrees   `json:"addrees" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Day         string    `json:"day" gorm:"type:varchar(255)"`
	Date        string    `json:"date" gorm:"type:varchar(255)"`
	Status      string    `json:"status" gorm:"type:varchar(255)"`
	Cart        []Cart    `json:"cart" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
	TotalAmount int       `json:"total_amount" gorm:"type:int"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type TransactionUserResponse struct {
	ID         int    `json:"id"`
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
	UserID     int    `json:"-"`
}

func (TransactionUserResponse) TableName() string {
	return "transaction"
}
