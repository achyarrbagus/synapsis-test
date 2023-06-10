package repostitories

import (
	"synapsis-test/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetTransactionActive(UserId int) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransactionActive(UserId int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Address").Preload("Address.User").Preload("Cart").Where("status = ? AND user_id = ?", "active", UserId).First(&transaction).Error
	return transaction, err
}
