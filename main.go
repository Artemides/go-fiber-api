package main

import (
	"log"

	"github.com/Artemides/go-fiber-api/initializers"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	app := fiber.New()

	app.Get("/api/greetings", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "alles ok",
			"message": "Welcome",
		})
	})

	log.Fatal(app.Listen(":4000"))
}
