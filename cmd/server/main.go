package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"

	"go-fiber-template/client/pgx"
	"go-fiber-template/client/valkey"
	"go-fiber-template/internal/config"
	"go-fiber-template/internal/server/route"
)

func main() {
	cfg, err := config.New(config.NewSetting())
	if err != nil {
		log.Fatal(err)
	}

	if client, err := valkey.NewClient(context.Background(), &valkey.Config{
		Address: cfg.Setting().ValkeyAddress,
		DB:      cfg.Setting().ValkeyDB,
	}); err != nil {
		log.Printf("[warn] valkey init failed: %v", err)
	} else {
		client.Close()
	}

	if db, err := pgx.NewDB(context.Background(), pgx.DBConfig{
		Host:     cfg.Setting().PostgresHost,
		Port:     cfg.Setting().PostgresPort,
		User:     cfg.Setting().PostgresUser,
		Password: cfg.Setting().PostgresPassword,
		Database: cfg.Setting().PostgresDatabase,
	}); err != nil {
		log.Printf("[warn] pgx init failed: %v", err)
	} else {
		db.Close()
	}

	app := fiber.New()
	route.Setup(app)

	if err := app.Listen(fmt.Sprintf(":%s", cfg.Setting().Port)); err != nil {
		log.Fatal(err)
	}
}
