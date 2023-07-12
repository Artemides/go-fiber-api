package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UploadSingleFileController(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	files := form.File["image"]

	for _, file := range files {
		fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
	}

	return c.SendStatus(fiber.StatusOK)

}
