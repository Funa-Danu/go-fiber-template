package v1

import (
	"github.com/gofiber/fiber/v3"

	"go-fiber-template/internal/server/handler"
)

// Register adds /v1 handlers.
func Register(api fiber.Router) {
	api.Get("/health", handler.Health)
}
