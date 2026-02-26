package main

import (
	"log"

	"github.com/gofiber/fiber/v3"

	"go-fiber-template/internal/router"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
