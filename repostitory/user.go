package repostitories

import (
	"synapsis-test/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user models.User) (models.User, error)
	GetUser(UserId int) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

// func (r *repository) GetUser(UserId int) (models.User, error) {
// 	var User models.User
// 	err := r.db.Raw("SELECT * FROM users WHERE id=?", UserId).Scan(&User).Error
// 	return User, err
// }

func (r *repository) GetUser(UserId int) (models.User, error) {
	var User models.User
	err := r.db.Preload("Transaction").First(&User, UserId).Error
	return User, err
}
