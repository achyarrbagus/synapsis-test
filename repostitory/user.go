package repostitories

import (
	"synapsis-test/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user models.User) (models.User, error)
	GetUser(UserId int) (models.User, error)
	Login(email string) (models.User, error)
	CheckAuth(ID int) (models.User, error)
	GetAllUser() ([]models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllUser() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Transaction").Preload("Address").Find(&users).Error

	return users, err
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}
func (r *repository) Login(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repository) GetUser(UserId int) (models.User, error) {
	var User models.User
	err := r.db.First(&User, UserId).Error
	return User, err
}

func (r *repository) CheckAuth(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Transaction").Preload("Address").First(&user, ID).Error

	return user, err
}
