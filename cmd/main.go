package main

import (
	userRoutes "github.com/OmeChat/server/internal/webserver/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(logger.New())
	app.Use(cors.New())

	userAPI := app.Group("/user-api")
	userAPI.Post("/create-account", userRoutes.CreateAccount)
	userAPI.Post("/add-client", userRoutes.AddClient)

	app.Listen(":8080")

}
