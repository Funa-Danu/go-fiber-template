package main

import (
	"log"

	"go-fiber-template/internal/server"
)

func main() {
	app := server.New()

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
