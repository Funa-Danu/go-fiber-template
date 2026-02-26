package server

import (
	"github.com/gofiber/fiber/v3"

	"go-fiber-template/internal/server/route"
)

// New creates and wires the Fiber app.
func New() *fiber.App {
	app := fiber.New()

	route.Setup(app)

	return app
}
