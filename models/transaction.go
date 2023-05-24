package models

type Transaction struct {
	ID          int    `json:"id" gorm:"primary_key:auto_increment"`
	UserID      int    `json:"-"`
	User        User   `json:"-" gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	Name        string `json:"name" gorm:"type:varchar(255)"`
	Address     string `json:"address" gorm:"type:varchar(255)"`
	PostalCode  string `json:"postal_code" gorm:"type:varchar(255)"`
	Phone       string `json:"phone" gorm:"type:varchar(255)"`
	Day         string `json:"day" gorm:"type:varchar(255)"`
	Date        string `json:"date" gorm:"type:varchar(255)"`
	Status      string `json:"status" gorm:"type:varchar(255)"`
	Cart        []Cart `json:"cart" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
	TotalAmount int    `json:"total_amount" gorm:"type:int"`
}
