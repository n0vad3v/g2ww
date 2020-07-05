package main

import (
	"os"

	"github.com/gofiber/compression"
	"github.com/gofiber/fiber"
)

func main() {

	app := fiber.New()

	var ListenAddress string
	if os.Getenv("DOCKER") != "" {
		ListenAddress = "0.0.0.0:2408"
	} else {
		ListenAddress = "127.0.0.1:2408"
	}

	// Server Info
	app.Use(compression.New())
	app.Get("/", GwStat())
	app.All("/:key", GwWorker())

	app.Listen(ListenAddress)

}
