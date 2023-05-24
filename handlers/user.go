package handlers

import (
	"net/http"
	"strconv"
	dto "synapsis-test/dto/result"
	userdto "synapsis-test/dto/user"
	"synapsis-test/models"
	"synapsis-test/pkg/bcrypt"
	repostitories "synapsis-test/repostitory"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type handlerUser struct {
	UserRepository repostitories.UserRepository
}

func HandlerUser(UserRepository repostitories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) Register(c echo.Context) error {
	request := new(userdto.UserRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	requestPassword := c.FormValue("password")
	password, err := bcrypt.HashingPassword(requestPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	requestName := c.FormValue("name")
	requestEmail := c.FormValue("email")

	user := models.User{
		Name:     requestName,
		Email:    requestEmail,
		Password: password,
		Role:     "user",
	}

	newData, err := h.UserRepository.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	data, _ := h.UserRepository.GetUser(newData.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlerUser) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: user})
}
