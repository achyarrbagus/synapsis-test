package routes

import (
	"synapsis-test/handlers"
	"synapsis-test/pkg/postgree"
	repostitories "synapsis-test/repostitory"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repostitories.RepositoryUser(postgree.DB)

	h := handlers.HandlerUser(userRepository)

	e.POST("/register", h.Register)
	e.GET("/user/:id", h.GetUser)

}
