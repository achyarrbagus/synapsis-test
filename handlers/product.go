package handlers

import (
	"net/http"
	"strconv"
	productdto "synapsis-test/dto/product"
	dto "synapsis-test/dto/result"
	"synapsis-test/models"
	repostitories "synapsis-test/repostitory"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type handlerProduct struct {
	ProductRepository  repostitories.ProductRepository
	CategoryRepository repostitories.CategoryRepository
}

func HandlerProduct(ProductRepository repostitories.ProductRepository, CategoryRepository repostitories.CategoryRepository) *handlerProduct {
	return &handlerProduct{ProductRepository, CategoryRepository}
}

func (h *handlerProduct) GetAllProduct(c echo.Context) error {
	product, err := h.ProductRepository.GetAllProduct()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: product})

}

func (h *handlerProduct) CreateProduct(c echo.Context) error {
	request := new(productdto.ProductRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	// for _, x := range request.CategoryID {
	// 	if int(x-0) > 0 {
	// 		getCategory, err := h.CategoryRepository.GetCategory(x)
	// 		if err != nil {
	// 			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Category Not Found"})
	// 		}
	// 		catagoryId = append(catagoryId, getCategory.ID)
	// 	}
	// }

	product := models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		Image:       request.Image,
		CategoryID:  request.CategoryID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	product, err = h.ProductRepository.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// product, _ = h.ProductRepository.GetProduct(product.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: product})

}

func (h *handlerProduct) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})

	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertProduct(data)})
}

func convertProduct(u models.Product) productdto.ProductResponse {
	return productdto.ProductResponse{
		Name:        u.Name,
		Description: u.Description,
		Price:       u.Price,
		Stock:       u.Stock,
		Image:       u.Image,
		CategoryID:  u.CategoryID,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}
