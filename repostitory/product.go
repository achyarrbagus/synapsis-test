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
	UpdateProduct(product models.Product) (models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) UpdateProduct(product models.Product) (models.Product, error) {
	r.db.Model(&product).Association("Category").Replace(product.Category)
	// .Association("Category"): Ini adalah metode dari objek model yang dikonfigurasi untuk mengakses relasi "Category". Melalui metode ini, kita dapat memanipulasi relasi antara objek product dan Category
	// Metode Replace digunakan untuk mengganti relasi yang ada dengan objek Category baru yang diberikan sebagai argumen (product.Category). Dengan kata lain, metode ini akan memperbarui relasi "Category" dari objek product dengan Category baru yang diberikan.
	err := r.db.Save(&product).Error
	return product, err
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
