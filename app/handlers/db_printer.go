package handlers

import (
	"github.com/drshooby/secrets-manager-practice/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DBPrintHandler(c echo.Context, db *gorm.DB) error {
	name := c.QueryParam("name")

	if name == "" {
		return c.JSON(300, echo.Map{"message": "name is empty"})
	}

	var songs []models.Song
	result := db.Where(&models.Song{Artist: name}).Find(&songs)
	if result.Error != nil {
		return c.JSON(500, echo.Map{"error": result.Error})
	}

	return c.JSON(200, songs)
}
