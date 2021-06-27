package routes

import (
	"github.com/gofiber/fiber/v2"
)

type defaultResponse struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

// The DefaultEndpoint is the endpoint that is being
// triggered if you try to access /
func DefaultEndpoint(ctx *fiber.Ctx) error {
	return ctx.JSON(defaultResponse{
		Message: "This is the official OmeChat backend",
		Version: "v1.0.0",
	})
}
