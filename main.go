package main

import (
	"log"

	"github.com/Artemides/go-fiber-api/controllers"
	"github.com/Artemides/go-fiber-api/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	micro.Route("/notes", func(router fiber.Router) {
		router.Get("", controllers.FindNotes)
		router.Post("/", controllers.CreateNoteHandler)
	})

	micro.Route("/note/:noteId", func(router fiber.Router) {
		router.Get("", controllers.FindNote)
		router.Patch("", controllers.UpdateNote)
		router.Delete("", controllers.DeleteNote)
	})

	micro.Get("/api/greetings", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "alles ok",
			"message": "Welcome",
		})
	})
	log.Fatal(app.Listen(":4000"))
}
