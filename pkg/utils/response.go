package utils

import "github.com/gofiber/fiber/v3"

// SuccessJSON is a shared success response helper.
func SuccessJSON(c fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{"ok": true, "data": data})
}
