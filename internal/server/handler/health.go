package handler

import "github.com/gofiber/fiber/v3"

func Root(c fiber.Ctx) error {
	return c.SendString("hello fiber3")
}

func Health(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"ok":      true,
		"service": "fiber3",
		"route":   "/v1/health",
	})
}
