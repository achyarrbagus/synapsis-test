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

	h := handlers.HandlerCart(CartRepository)

	e.POST("/cart", middleware.Auth(h.CreateCart))

}
