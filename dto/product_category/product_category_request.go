package dtoproductcategory

type ProductCategoryRequest struct {
	ID         int `json:"id" gorm:"primary_key:auto_increment"`
	ProductID  int `json:"-" form:"product_id" gorm:"foreignkey:ProductID"`
	CategoryID int `json:"-" form:"category_id" gorm:"foreignkey:CategoryID"`
}
