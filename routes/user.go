package routes

import (
	"synapsis-test/handlers"
	"synapsis-test/pkg/middleware"
	"synapsis-test/pkg/postgree"
	repostitories "synapsis-test/repostitory"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repostitories.RepositoryUser(postgree.DB)

	h := handlers.HandlerUser(userRepository)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("/user/:id", h.GetUser)
	e.GET("/user", h.GetAllUser)
	e.GET("/check-auth", middleware.Auth(h.CheckAuth))

}
