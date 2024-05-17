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

	routes.InitRouter(e, dbMysql)
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":8080"))
}
