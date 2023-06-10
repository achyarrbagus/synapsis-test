package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	UserRoutes(e)
	CategoryRoutes(e)
	CartRoutes(e)
	ProductRoutes(e)
	CartRoutes(e)
	AdreesRoutes(e)

}
