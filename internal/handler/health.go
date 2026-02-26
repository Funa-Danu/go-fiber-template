package handler

import "github.com/gofiber/fiber/v3"

// Root handles the root endpoint.
func Root(c fiber.Ctx) error {
	return c.SendString("hello fiber3")
}
