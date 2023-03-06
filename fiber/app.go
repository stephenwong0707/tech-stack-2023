package main

import (
	"store/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	handler.Init()
	v1 := app.Group("/api/v1")
	v1.Post("/message", handler.PostMessage)

	app.Listen(":8082")
}
