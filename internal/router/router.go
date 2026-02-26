package router

import (
	"github.com/gofiber/fiber/v3"

	"go-fiber-template/internal/handler"
)

// SetupRoutes wires application routes.
func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.Root)
}
