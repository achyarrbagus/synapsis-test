package handlers

import (
	"net/http"
	"strconv"
	categorydto "synapsis-test/dto/category"
	dto "synapsis-test/dto/result"
	"synapsis-test/models"
	repostitories "synapsis-test/repostitory"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type handlerCategory struct {
	CategoryRepository repostitories.CategoryRepository
}

func HandlerCategory(CategoryRepository repostitories.CategoryRepository) *handlerCategory {
	return &handlerCategory{CategoryRepository}
}

func (h *handlerCategory) CreateCategory(c echo.Context) error {
	request := new(categorydto.CategoryRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data := models.Category{
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newData, err := h.CategoryRepository.CreateCategory(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: newData})

}

func (h *handlerCategory) DeleteCategory(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	data, err := h.CategoryRepository.GetCategory(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})

	}
	delData, err := h.CategoryRepository.DeleteCategory(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})

	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: delData})

}

func (h *handlerCategory) GetAllCategory(c echo.Context) error {
	data, err := h.CategoryRepository.GetAllCategory()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlerCategory) GetCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlerCategory) UpdatedCategory(c echo.Context) error {
	request := new(categorydto.CategoryRequest)

	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	category, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		category.Name = request.Name
	}
	category.UpdatedAt = time.Now()

	data, err := h.CategoryRepository.UpdateCategory(category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})

}
