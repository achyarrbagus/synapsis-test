package routes

import (
	"synapsis-test/handlers"
	"synapsis-test/pkg/postgree"
	repostitories "synapsis-test/repostitory"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repostitories.RepositoryProduct(postgree.DB)
	categoryRepository := repostitories.RepositoryCategory(postgree.DB)
	productCategoryRepository := repostitories.RepositoryProductCategory(postgree.DB)

	h := handlers.HandlerProduct(productRepository, categoryRepository, productCategoryRepository)

	e.POST("/product", h.CreateProduct)
	e.GET("/product/:id", h.GetProduct)
	e.GET("/product", h.GetAllProduct)
	e.PATCH("/product/:id", h.UpdateProduct)
	e.DELETE("/product/:id", h.DeleteProduct)

}
