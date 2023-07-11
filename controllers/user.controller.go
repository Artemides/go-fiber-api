package controllers

import (
	"github.com/Artemides/go-fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": user})
}
