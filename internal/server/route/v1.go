package route

import (
	"github.com/gofiber/fiber/v3"

	"go-fiber-template/internal/server/handler"
	routev1 "go-fiber-template/internal/server/route/v1"
)

// Setup wires top-level routes and route groups.
func Setup(app *fiber.App) {
	api := app.Group("/")

	api.Get("", handler.Root)

	v1 := api.Group("/v1")
	routev1.Register(v1)
}
