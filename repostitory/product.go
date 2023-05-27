package repostitories

import (
	"synapsis-test/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(Product models.Product) (models.Product, error)
	GetProduct(Id int) (models.Product, error)
	DeleteProduct(product models.Product) (models.Product, error)
	GetAllProduct() ([]models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateProduct(Product models.Product) (models.Product, error) {
	err := r.db.Create(&Product).Error

	return Product, err
}

func (r *repository) GetAllProduct() ([]models.Product, error) {
	var product []models.Product
	err := r.db.Preload("Category").Find(&product).Error

	return product, err

}

func (r *repository) GetProduct(Id int) (models.Product, error) {
	var product models.Product
	err := r.db.Preload("Category").First(&product, Id).Error

	return product, err
}

func (r *repository) DeleteProduct(product models.Product) (models.Product, error) {
	err := r.db.Delete(&product).Error

	return product, err
}
