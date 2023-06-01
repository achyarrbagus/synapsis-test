package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	dtoaddrees "synapsis-test/dto/addrees"
	dto "synapsis-test/dto/result"
	"synapsis-test/models"
	repostitories "synapsis-test/repostitory"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerAddrees struct {
	UserRepository    repostitories.UserRepository
	AddreesRepository repostitories.AddreesRepository
}

func HandlerAddress(UserRepository repostitories.UserRepository, AddreesRepository repostitories.AddreesRepository) *handlerAddrees {
	return &handlerAddrees{UserRepository, AddreesRepository}
}

func (h *handlerAddrees) UpdateAddreesUser(c echo.Context) error {
	request := new(dtoaddrees.UpdateAddreesRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	// get user from context
	userLogin := c.Get("userLogin").(jwt.MapClaims)["id"].(float64)

	allUserAddrees, err := h.AddreesRepository.GetAllUserAddrees(int(userLogin))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	// validation struct request
	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	// get id from param
	id, _ := strconv.Atoi(c.Param("id"))
	addreesId := 0
	for _, item := range allUserAddrees {
		if item.ID == id {
			addreesId = item.ID
			break
		}
	}
	if addreesId == 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: 404, Message: "User Addrees Not Found"})
	}

	addrees, err := h.AddreesRepository.GetAdressById(addreesId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})

	}

	if request.FullName != "" {
		addrees.FullName = request.FullName
	}
	if request.City != "" {
		addrees.City = request.City
	}
	if request.HouseNumber != "" {
		addrees.HouseNumber = request.HouseNumber
	}
	if request.PostCode != "" {
		addrees.PostCode = request.PostCode
	}
	if request.Street != "" {
		addrees.Street = request.Street
	}
	if request.Province != "" {
		addrees.Province = request.Province
	}
	addrees.UserID = int(userLogin)
	user, err := h.UserRepository.GetUser(addrees.UserID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	addrees.User = user

	h.AddreesRepository.UpdateAddrees(addrees)
	return c.JSON(http.StatusBadRequest, dto.SuccessResult{Code: http.StatusOK, Data: addrees})
}

func (h *handlerAddrees) CreateAddrees(c echo.Context) error {
	request := new(dtoaddrees.AddreesRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// get user id from context
	loginUser := c.Get("userLogin")
	isLoginUser := loginUser.(jwt.MapClaims)["id"].(float64)
	// get user data from user_id
	user, _ := h.UserRepository.GetUser(int(isLoginUser))

	reqData := models.Addrees{
		FullName:    request.FullName,
		Street:      request.Street,
		HouseNumber: request.HouseNumber,
		PostCode:    request.PostCode,
		City:        request.City,
		Province:    request.Province,
		Country:     request.Country,
		UserID:      int(isLoginUser),
		User:        user,
	}
	reqData, err = h.AddreesRepository.CreateAddrees(reqData)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})

	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: reqData})

}

func (h *handlerAddrees) GetAllUserAddrees(c echo.Context) error {
	// get user id from context
	loginUser := c.Get("userLogin")
	isLoginUser := loginUser.(jwt.MapClaims)["id"].(float64)
	// get user data from user_id
	addressUser, err := h.AddreesRepository.GetAllUserAddrees(int(isLoginUser))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})

	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: addressUser})
}

func (h *handlerAddrees) GetOneUserAddrees(c echo.Context) error {
	addreesId, _ := strconv.Atoi(c.Param("id"))

	// Get user Id from context
	loginUser := c.Get("userLogin")
	isLoginUser := loginUser.(jwt.MapClaims)["id"].(float64)

	// get one user Addrees
	data, err := h.AddreesRepository.GetAllUserAddrees(int(isLoginUser))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	var userAddrees models.Addrees
	for _, x := range data {
		if x.ID == addreesId {
			userAddrees = x
		}

	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: userAddrees})
}

func (h *handlerAddrees) DeleteUserAddreess(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println(id)

	// Get user Id from context
	loginUser := c.Get("userLogin")
	isLoginUser := loginUser.(jwt.MapClaims)["id"].(float64)

	data, err := h.AddreesRepository.GetAllUserAddrees(int(isLoginUser))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	var userAddrees models.Addrees
	for _, x := range data {
		if x.ID == id {
			userAddrees = x
		}

	}

	fmt.Println(userAddrees)

	delData, err := h.AddreesRepository.DeleteUserAddrees(userAddrees)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: delData})

}
func (h *handlerAddrees) GetAdressById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := h.AddreesRepository.GetAdressById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: 404, Message: "Addrees Not Found"})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})

}
