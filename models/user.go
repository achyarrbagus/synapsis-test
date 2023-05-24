package models

import "time"

type User struct {
	ID          int           `json:"id" gorm:"primary_key:auto_increment"`
	Name        string        `json:"fullname" form:"fullname" gorm:"type:varchar(255)"`
	Email       string        `json:"email" form:"email" gorm:"type:varchar(255)"`
	Password    string        `json:"password" form:"password" gorm:"type:varchar(255)"`
	Transaction []Transaction `json:"transaction"`
	Role        string        `json:"role"`
	CreatedAt   time.Time     `json:"-"`
	UpdatedAt   time.Time     `json:"-"`
}

type UsersTransactionResponse struct {
	ID          int           `json:"id"  gorm:"primary_key:auto_increment"`
	Name        string        `json:"fullname" form:"fullname" gorm:"type: varchar(255)"`
	Transaction []Transaction `json:"transaction"`
	Role        string        `json:"role"`
	CreatedAt   time.Time     `json:"-"`
	UpdatedAt   time.Time     `json:"-"`
}

type UsersResponse struct {
	ID        int       `json:"id"  gorm:"primary_key:auto_increment"`
	Name      string    `json:"fullname" form:"fullname" gorm:"type: varchar(255)"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (UsersTransactionResponse) TableName() string {
	return "users"
}

func (UsersResponse) TableName() string {
	return "users"
}
