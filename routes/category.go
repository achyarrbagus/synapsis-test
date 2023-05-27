package routes

import (
	"synapsis-test/handlers"
	"synapsis-test/pkg/postgree"
	repostitories "synapsis-test/repostitory"

	"github.com/labstack/echo/v4"
)

func categoryRoutes(e *echo.Group) {
	CategoryRepository := repostitories.RepositoryCategory(postgree.DB)
	h := handlers.HandlerCategory(CategoryRepository)

	e.POST("/category", h.CreateCategory)
	e.GET("/category", h.GetAllCategory)
	e.GET("/category/:id", h.GetCategory)
	e.DELETE("/category/:id", h.DeleteCategory)
	e.PATCH("/category/:id", h.UpdatedCategory)
}
