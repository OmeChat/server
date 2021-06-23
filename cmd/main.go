package main

import (
	userRoutes "github.com/OmeChat/server/internal/webserver/routes/user"
	ws "github.com/OmeChat/server/internal/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Use(logger.New())
	app.Use(cors.New())

	app.Use("/ws", websocket.New(ws.Router))

	ws.WS_DATAFLOW_CHANNEL = *new(chan ws.ConnectionPair)
	go ws.DataHandler()

	userAPI := app.Group("/user-api")
	userAPI.Post("/create-account", userRoutes.CreateAccount)
	userAPI.Post("/add-client", userRoutes.AddClient)

	app.Listen(":8080")

}
