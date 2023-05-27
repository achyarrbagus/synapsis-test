package models

type Cart struct {
	ID            int         `json:"id" gorm:"primary_key:auto_increment"`
	UserID        int         `json:"user_id"`
	User          User        `json:"user" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;foreignKey:UserID"`
	ProductID     int         `json:"-"`
	Product       Product     `json:"-" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;foreignKey:ProductID"`
	Qty           int         `json:"qty" gorm:"type:int"`
	TransactionID int         `json:"transaction_id" gorm:"type: int"`
	Transaction   Transaction `json:"-"`
}
