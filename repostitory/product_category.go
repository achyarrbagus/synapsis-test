package repostitories

import (
	"synapsis-test/models"

	"gorm.io/gorm"
)

type ProductCategory interface {
	CreateProductCategory(Product models.Product) (models.Product, error)
}

func RepositoryProductCategory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateProductCategory(Product models.Product) (models.Product, error) {
	err := r.db.Create(&Product).Error

	return Product, err
}
