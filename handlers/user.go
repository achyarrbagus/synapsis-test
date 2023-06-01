package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	dto "synapsis-test/dto/result"
	userdto "synapsis-test/dto/user"
	"synapsis-test/models"
	"synapsis-test/pkg/bcrypt"
	jwtToken "synapsis-test/pkg/jwt"
	repostitories "synapsis-test/repostitory"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
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

	requestPassword := request.Password
	password, err := bcrypt.HashingPassword(requestPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
		Role:     "user",
	}

	newData, err := h.UserRepository.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: newData})
}

func (h *handlerUser) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: user})
}

func (h *handlerUser) Login(c echo.Context) error {
	request := new(userdto.LoginRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	if err := validation.Struct(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user, err := h.UserRepository.Login(request.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "email or password is wrong"})
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	fmt.Println(claims)

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginResponse := userdto.LoginResponse{
		Email: user.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: loginResponse})
}

func (h *handlerUser) CheckAuth(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, _ := h.UserRepository.CheckAuth(int(userId))

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: user})
}

func (h *handlerUser) GetAllUser(c echo.Context) error {
	users, err := h.UserRepository.GetAllUser()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: users})
}
