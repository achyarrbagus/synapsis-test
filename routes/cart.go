package routes

import (
	"synapsis-test/handlers"
	"synapsis-test/pkg/middleware"
	"synapsis-test/pkg/postgree"
	repostitories "synapsis-test/repostitory"

	"github.com/labstack/echo/v4"
)

func CartRoutes(e *echo.Group) {
	CartRepository := repostitories.RepositoryCart(postgree.DB)
	TransactionRepository := repostitories.RepositoryTransaction(postgree.DB)
	UserRepository := repostitories.RepositoryUser(postgree.DB)
	AddreesRepository := repostitories.RepositoryAddress(postgree.DB)
	ProductRepository := repostitories.RepositoryProduct(postgree.DB)

	h := handlers.HandlerCart(CartRepository, TransactionRepository, UserRepository, AddreesRepository, ProductRepository)

	e.POST("/cart", middleware.Auth(h.CreateCart))

}
