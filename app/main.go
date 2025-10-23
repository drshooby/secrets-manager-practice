package main

import (
	"github.com/drshooby/secrets-manager-practice/handlers"
	"github.com/drshooby/secrets-manager-practice/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	dsn := "david:password@tcp(db:3306)/test-db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Song{})

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hi from Echo microservice!")
	})

	e.GET("/artist", func(c echo.Context) error {
		return handlers.DBPrintHandler(c, db)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
