package productdto

import (
	"synapsis-test/models"
	"time"
)

type ProductResponse struct {
	ID          int               `json:"id" gorm:"primary_key:auto_increment"`
	Name        string            `json:"name" form:"name" gorm:"type:varchar(255)"`
	Description string            `json:"description" form:"description" gorm:"type:varchar(255)"`
	Price       int               `json:"price" form:"price" gorm:"type:int"`
	Stock       int               `json:"stock" form:"stock" gorm:"type:int"`
	Image       string            `json:"image" form:"image" gorm:"type:varchar(255)"`
	Category    []models.Category `json:"-"`
	CategoryID  []int             `json:"category_id" form:"category_id" validate:"required"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}
