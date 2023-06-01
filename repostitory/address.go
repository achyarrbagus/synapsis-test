package repostitories

import (
	"synapsis-test/models"

	"gorm.io/gorm"
)

type AddreesRepository interface {
	CreateAddrees(address models.Addrees) (models.Addrees, error)
	GetOneUserAddress(UserAddreesId int, UserId int) (models.Addrees, error)
	GetAllUserAddrees(UserID int) ([]models.Addrees, error)
	GetAdressById(AddreesID int) (models.Addrees, error)
	DeleteUserAddrees(UserAddreesId models.Addrees) (models.Addrees, error)
	UpdateAddrees(addrees models.Addrees) (models.Addrees, error)
}

func RepositoryAddress(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateAddrees(address models.Addrees) (models.Addrees, error) {
	err := r.db.Create(&address).Error
	return address, err
}

func (r *repository) UpdateAddrees(addrees models.Addrees) (models.Addrees, error) {
	err := r.db.Save(&addrees).Error
	return addrees, err
}

func (r *repository) GetAllUserAddrees(UserID int) ([]models.Addrees, error) {
	var userAddress []models.Addrees
	err := r.db.Preload("User").Where("user_id", UserID).Find(&userAddress).Error
	return userAddress, err
}

func (r *repository) GetAdressById(AddreesID int) (models.Addrees, error) {
	var Addrees models.Addrees
	err := r.db.Preload("User").First(&Addrees, AddreesID).Error
	return Addrees, err
}

func (r *repository) GetOneUserAddress(UserAddreesId int, UserId int) (models.Addrees, error) {
	var userAddrees models.Addrees

	err := r.db.Raw("SELECT * FROM addrees WHERE id=? AND user_id=?", UserAddreesId, UserId).Error

	return userAddrees, err

}

func (r *repository) DeleteUserAddrees(UserAddreesId models.Addrees) (models.Addrees, error) {
	err := r.db.Delete(&UserAddreesId).Error

	return UserAddreesId, err
}
