package main

import (
	"log"

	"github.com/gofiber/fiber/v3"

	"go-fiber-template/internal/server/route"
)

func main() {
	app := fiber.New()
	route.Setup(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
