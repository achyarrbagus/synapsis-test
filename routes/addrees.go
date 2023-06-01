package routes

import (
	"synapsis-test/handlers"
	"synapsis-test/pkg/middleware"
	"synapsis-test/pkg/postgree"
	repostitories "synapsis-test/repostitory"

	"github.com/labstack/echo/v4"
)

func AdreesRoutes(e *echo.Group) {
	addreesRepository := repostitories.RepositoryAddress(postgree.DB)
	userRepository := repostitories.RepositoryUser(postgree.DB)

	h := handlers.HandlerAddress(userRepository, addreesRepository)

	e.POST("/addrees", middleware.Auth(h.CreateAddrees))
	e.GET("/addrees/:id", h.GetAdressById)
	e.GET("/addrees/users", middleware.Auth(h.GetAllUserAddrees))
	e.GET("/addrees/user", middleware.Auth(h.GetOneUserAddrees))
	e.PATCH("/addrees/user/:id", middleware.Auth(h.UpdateAddreesUser))
	e.DELETE("/addrees/user/:id", middleware.Auth(h.DeleteUserAddreess))

}
