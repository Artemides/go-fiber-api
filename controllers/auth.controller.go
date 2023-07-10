package controllers

import (
	"strings"

	"github.com/Artemides/go-fiber-api/initializers"
	"github.com/Artemides/go-fiber-api/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {
	var payload *models.SignUpInput
	err := c.BodyParser(&payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error parsing input", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": errors})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed hashing password", "message": err.Error()})
	}

	newUser := models.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: string(hashedPassword),
		Photo:    &payload.Photo,
	}
	response := initializers.DB.Create(&newUser)

	if response.Error != nil && strings.Contains(response.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": response.Error.Error()})
	} else if response.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "creation failed", "message": response.Error.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": fiber.Map{"user": models.FilterUserRecord(&newUser)}})
}
