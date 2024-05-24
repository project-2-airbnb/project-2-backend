package main

import (
	"project-2/app/config"
	"project-2/app/databases"
	"project-2/app/migrations"
	"project-2/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	dbMysql := databases.InitDBMysql(cfg)
	migrations.RunMigrations(dbMysql)

	e := echo.New()

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	// CORS configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Ganti "*" dengan origin yang lebih spesifik jika perlu
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Routes
	routes.InitRouter(e, dbMysql)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

