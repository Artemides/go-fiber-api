package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/Artemides/go-fiber-api/initializers"
	"github.com/Artemides/go-fiber-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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

func SignInController(c *fiber.Ctx) error {
	var payload *models.SignInInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": errors})
	}

	var user *models.User

	response := initializers.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if err := response.Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})

	}

	config, _ := initializers.LoadConfig(".")
	tokenBytes := jwt.New(jwt.SigningMethodHS256)
	now := time.Now().UTC()
	claims := tokenBytes.Claims.(jwt.MapClaims)

	claims["sub"] = user.ID
	claims["exp"] = now.Add(config.JwtExpiredIn).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Add(config.JwtExpiredIn).Unix()

	tokenString, err := tokenBytes.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "failed", "message": fmt.Sprintf("generate JWT token failed %v", err.Error())})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   config.JwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": tokenString})
}
