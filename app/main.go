package main

import (
	builder "github.com/drshooby/secrets-manager-practice/aws"
	"github.com/drshooby/secrets-manager-practice/handlers"
	"github.com/drshooby/secrets-manager-practice/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	awsConfig := builder.AWSConfigLoader("us-east-1")
	dsn := builder.BuildDSN(awsConfig, "a14db")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Song{})

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hi from Echo microservice!")
	})

	e.GET("/artist", func(c echo.Context) error {
		return handlers.DBPrintHandler(c, db, awsConfig)
	})

	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(":8080"))
}
