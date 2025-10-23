package main

import (
	"github.com/drshooby/secrets-manager-practice/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.DBPrintHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
