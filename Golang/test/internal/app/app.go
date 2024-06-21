package app

import (
	"test/internal/config"
	"test/internal/handlers"
	"test/internal/middleware"
	"test/pkg/db"
	"test/pkg/logger"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Start() {
	db.InitDB()
	app := fiber.New()

	app.Post("/login", handlers.Login)

	app.Use(middleware.AuthMiddleware)

	app.Post("/edit/:Id", handlers.EditNews)
	app.Get("/list", handlers.ListNews)

	logger.Init()
	//app.Listen(":" + config.AppConfig.Port)

	if err := app.Listen(":" + config.AppConfig.Port); err != nil {
		log.Fatalf("Ошибка запуска сервера: "+ err.Error())
	}
}

