package repostitories

import (
	"synapsis-test/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	CreateCart(cart models.Cart) (models.Cart, error)
	GetAllCart() ([]models.Cart, error)
	GetCart(ID int) (models.Cart, error)
	GetAllUserCart(UserID int) ([]models.Cart, error)
	GetOneUserCart(UserID int, ProductID int) (models.Cart, error)
	UpdateOneCartUser(Cart models.Cart) (models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error
	return cart, err
}

func (r *repository) GetAllCart() ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Find(&cart).Error
	return cart, err
}

func (r *repository) GetAllUserCart(UserID int) ([]models.Cart, error) {
	var userCart []models.Cart
	err := r.db.Preload("User").Where("user_id", UserID).Find(&userCart).Error
	return userCart, err
}

func (r *repository) UpdateOneCartUser(Cart models.Cart) (models.Cart, error) {
	err := r.db.Save(&Cart).Error
	return Cart, err
}

func (r *repository) GetOneUserCart(UserID int, ProductID int) (models.Cart, error) {
	var productCart models.Cart
	err := r.db.Preload("Transaction").Preload("Transaction.User").Preload("Product").Preload("User").Where("user_id = ? AND product_id = ?", UserID, ProductID).First(&productCart).Error
	// err := r.db.First("SELECT * FROM product_carts WHERE user_id = ? AND product_id = ?", UserID, ProductID).Scan(&productCart).Error
	return productCart, err
}

func (r *repository) GetCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.First("SELECT * FROM cart WHERE cart_id=?", ID).Error
	return cart, err
}
func (r *repository) UpdateCart(cart models.Cart) (models.Cart, error) {
	// r.db.Model(&product).Association("Category").Replace(product.Category)
	// .Association("Category"): Ini adalah metode dari objek model yang dikonfigurasi untuk mengakses relasi "Category". Melalui metode ini, kita dapat memanipulasi relasi antara objek product dan Category
	// Metode Replace digunakan untuk mengganti relasi yang ada dengan objek Category baru yang diberikan sebagai argumen (product.Category). Dengan kata lain, metode ini akan memperbarui relasi "Category" dari objek product dengan Category baru yang diberikan.
	err := r.db.Save(&cart).Error
	return cart, err
}

// func (r *repository) GetAllCartUser(UserID int) ([]models.Cart, error) {
// 	var userCart []models.Cart
// 	err := r.db.Where("user_id = ?", UserID).Find(&userCart).Error

// 	return userCart, err
// }

// func (r *repository) GetPendingTransactionUser(UserID int) (models.Transaction, error) {
// 	var pendingTransUser models.Transaction
// 	err := r.db.Where("status = ? AND user_id = ?", "pending", UserID).First(&pendingTransUser).Error

// 	return pendingTransUser, err

// }
