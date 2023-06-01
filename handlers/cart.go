package handlers

import (
	"fmt"
	"net/http"
	cartdto "synapsis-test/dto/cart"
	dto "synapsis-test/dto/result"
	repostitories "synapsis-test/repostitory"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerCart struct {
	CartRepository repostitories.CartRepository
	// UserRepository    repostitories.UserRepository
	// ProductRepository repostitories.ProductRepository
}

func HandlerCart(CartRepository repostitories.CartRepository) *handlerCart {
	return &handlerCart{CartRepository}
}

func (h *handlerCart) CreateCart(c echo.Context) error {
	request := new(cartdto.CartRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	loginUser := c.Get("userLogin")
	isLoginUser := loginUser.(jwt.MapClaims)["id"].(float64)

	fmt.Println(request, "ini request")
	fmt.Println(request.Qty, "ini data qty")
	fmt.Println(isLoginUser, "ini user yang login")

	// getUser, err := h.UserRepository.GetUser(int(isLoginUser))
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "User Not Found"})

	// }
	// getProduct, err := h.ProductRepository.GetProduct(request.ProductID)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Product Not Found"})

	// }
	// fmt.Println(getProduct, "ini data product")
	// dataReq := models.Cart{
	// 	UserID:    int(isLoginUser),
	// 	User:      getUser,
	// 	ProductID: getProduct.ID,
	// 	Product:   getProduct,
	// 	Qty:       request.Qty,
	// }
	// fmt.Println(dataReq, "ini datareq")

	// data, _ := h.CartRepository.CreateCart(dataReq)

	// datas, _ := h.CartRepository.GetCart(data.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: isLoginUser})

}
