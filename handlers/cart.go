package handlers

import (
	"math/rand"
	"net/http"
	cartdto "synapsis-test/dto/cart"
	dto "synapsis-test/dto/result"
	"synapsis-test/models"
	repostitories "synapsis-test/repostitory"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerCart struct {
	CartRepository        repostitories.CartRepository
	UserRepository        repostitories.UserRepository
	TransactionRepository repostitories.TransactionRepository
	AddreesRepository     repostitories.AddreesRepository
	// TransactionRepository repostitories.TransactionRepository
	ProductRepository repostitories.ProductRepository
}

func HandlerCart(CartRepository repostitories.CartRepository, UserRepository repostitories.UserRepository, TransactionRepository repostitories.TransactionRepository, AddreesRepository repostitories.AddreesRepository, ProductRepository repostitories.ProductRepository) *handlerCart {
	return &handlerCart{CartRepository, UserRepository, TransactionRepository, AddreesRepository, ProductRepository}
}

func (h *handlerCart) CreateCart(c echo.Context) error {
	// get user from contex
	loginUser := c.Get("userLogin")
	isLoginUser := loginUser.(jwt.MapClaims)["id"].(float64)

	// get user
	user, _ := h.UserRepository.GetUser(int(isLoginUser))

	activeTrans, err := h.TransactionRepository.GetTransactionActive(int(isLoginUser))
	if err != nil {
		//to get random user addrees
		rand.Seed(time.Now().UnixNano())
		// get data user addrees from id
		// get random addrees from user
		dataAddrees, err := h.AddreesRepository.GetAllUserAddrees(int(isLoginUser))
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: 404, Message: err.Error()})
		}
		randomIndex := rand.Intn(len(dataAddrees))
		randomData := dataAddrees[randomIndex]

		newTrans := models.Transaction{
			ID:        int(time.Now().Unix()),
			UserID:    int(isLoginUser),
			User:      user,
			Status:    "active",
			AddressID: randomIndex,
			Address:   randomData,
		}

		data, err := h.TransactionRepository.CreateTransaction(newTrans)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: 404, Message: err.Error()})
		}

		request := new(cartdto.CartRequest)
		if err := c.Bind(request); err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}

		product, err := h.ProductRepository.GetProduct(request.ProductID)

		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: 404, Message: "Product Not Found"})
		}
		newCart := models.Cart{
			UserID:        int(isLoginUser),
			User:          user,
			ProductID:     product.ID,
			Product:       product,
			Qty:           1,
			Price:         product.Price,
			TransactionID: data.ID,
			Transaction:   data,
		}
		cartData, err := h.CartRepository.CreateCart(newCart)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}
		return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: cartData})
	}
	//
	//
	//
	request := new(cartdto.CartRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	productCart, err := h.CartRepository.GetOneUserCart(int(isLoginUser), request.ProductID)
	if err != nil {
		product, err := h.ProductRepository.GetProduct(request.ProductID)

		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: 404, Message: "Product Not Found"})
		}
		newCart := models.Cart{
			UserID:        int(isLoginUser),
			User:          user,
			ProductID:     product.ID,
			Product:       product,
			Qty:           1,
			Price:         product.Price,
			TransactionID: activeTrans.ID,
			Transaction:   activeTrans,
		}
		data, err := h.CartRepository.CreateCart(newCart)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}
		return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})

	}

	productCart.Qty = productCart.Qty + 1
	productCart.Price = productCart.Qty * productCart.Product.Price

	newData, err := h.CartRepository.UpdateOneCartUser(productCart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertCart(newData)})

}

func convertCart(u models.Cart) cartdto.CartResponse {
	return cartdto.CartResponse{
		Qty:           u.Qty,
		ID:            u.ID,
		UserID:        u.ID,
		User:          u.User,
		ProductID:     u.ProductID,
		Product:       u.Product,
		Price:         u.Price,
		TransactionID: u.TransactionID,
		Transaction:   u.Transaction,
	}
}
