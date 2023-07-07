package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)


func main()  {
	app:=fiber.New();

	app.Get("/api/health",func (c *fiber.Ctx) error  {
		return c.Status(200).JSON(fiber.Map{
			"status":"Alles Gut",
			"message":"Hezlich Willkommen to GO",
		})
	})

	log.Fatal(app.Listen(":4000"))
}