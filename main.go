package main

import (
	"fmt"
	"os"
	"synapsis-test/database"
	"synapsis-test/pkg/postgree"
	"synapsis-test/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	var port = os.Getenv("PORT")
	fmt.Println(port)

	// initialiation echo
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	postgree.DatabaseInit()
	database.RunMigration()
	routes.RouteInit(e.Group("/api/v1"))

	e.Static("/uploads", "./uploads")
	fmt.Println("server running localhost:" + port)
	e.Logger.Fatal(e.Start("localhost:" + port))
}
