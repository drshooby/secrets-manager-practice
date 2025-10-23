package handlers

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/drshooby/secrets-manager-practice/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DBPrintHandler(c echo.Context, db *gorm.DB, cfg aws.Config) error {
	name := c.QueryParam("name")

	if name == "" {
		return c.JSON(http.StatusNoContent, echo.Map{"message": "name is empty"})
	}

	var songs []models.Song
	result := db.Where(&models.Song{Artist: name}).Find(&songs)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error})
	}

	return c.JSON(http.StatusOK, songs)
}
