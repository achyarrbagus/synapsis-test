package cartdto

import "synapsis-test/models"

type CartResponse struct {
	ID            int                `json:"id" gorm:"primary_key:auto_increment"`
	UserID        int                `json:"user_id" form:"user_id" validate:"required"`
	User          models.User        `json:"user"`
	ProductID     int                `json:"product_id" form:"product_id" validate:"required"`
	Product       models.Product     `json:"product"`
	Qty           int                `json:"qty"`
	Price         int                `json:"price"`
	TransactionID int                `json:"transaction_id" gorm:"type: int"`
	Transaction   models.Transaction `json:"transaction"`
}
