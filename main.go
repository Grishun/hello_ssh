package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/test", func(ctx fiber.Ctx) error {
		return ctx.SendString("I'm alive")
	})

	log.Fatal(app.Listen(":8080"))
}
