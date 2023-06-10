package models

import "time"

type Addrees struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	FullName    string    `json:"full_name"`
	Street      string    `json:"street"`
	HouseNumber string    `json:"house_number"`
	Phone       string    `json:"phone" gorm:"type:varchar(255)"`
	PostCode    string    `json:"post_code"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	Country     string    `json:"country"`
	UserID      int       `json:"user_id" `
	User        User      `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
