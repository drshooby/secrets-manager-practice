package main

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/drshooby/secrets-manager-practice/handlers"
	"github.com/drshooby/secrets-manager-practice/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	sm "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func awsString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func main() {
	e := echo.New()

	secretName := "test"
	awsRegion := "us-east-1"
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(awsRegion))
	if err != nil {
		panic(err.Error())
	}

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

	e.GET("/secret", func(c echo.Context) error {
		client := sm.NewFromConfig(cfg)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		secretOutput, err := client.GetSecretValue(ctx, &sm.GetSecretValueInput{
			SecretId: &secretName,
		})
		if err != nil {
			return c.JSON(500, echo.Map{"error": err.Error()})
		}

		return c.JSON(200, echo.Map{
			"is this your secret?": awsString(secretOutput.SecretString),
		})
	})

	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(":8080"))
}
