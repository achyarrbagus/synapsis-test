package models

type Address struct {
	ID          int    `json:"-" gorm:"primary_key:auto_increment"`
	FullName    string `json:"-"`
	Street      string `json:"street"`
	HouseNumber string `json:"-"`
	PostCode    string `json:"-"`
	City        string `json:"-"`
	Province    string `json:"province"`
	Country     string `json:"country"`
	UserID      int    `json:"user_id"`
	User        User   `json:"-"`
}
