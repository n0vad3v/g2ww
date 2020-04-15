package main

import (
	"github.com/gofiber/compression"
	"github.com/gofiber/fiber"
)

func main() {

	app := fiber.New()

	ListenAddress := "127.0.0.1:2408"

	// Server Info
	app.Use(compression.New())
	app.All("/:key", GwWorker())

	app.Listen(ListenAddress)

}
