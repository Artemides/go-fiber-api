package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Artemides/go-fiber-api/initializers"
	"github.com/Artemides/go-fiber-api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateNoteHandler(c *fiber.Ctx) error {
	var payload *models.CreateNoteSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	now := time.Now()

	newNote := models.Note{
		Title:     payload.Title,
		Content:   payload.Content,
		Category:  payload.Category,
		Published: payload.Published,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := initializers.DB.Create(&newNote)

	if result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate entry") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "title already exists, please use another note title"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": newNote}})
}

func FindNotes(c *fiber.Ctx) error {
	page := c.Query("page", "1")
	limit := c.Query("limit", "10")
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var notes []models.Note
	response := initializers.DB.Limit(intLimit).Offset(offset).Find(&notes)

	if response.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"stauts": "error", "message": response.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(notes), "data": notes})
}

func FindNote(c *fiber.Ctx) error {
	noteId := c.Params("noteId")

	var note models.Note

	response := initializers.DB.First(&note, "id = ?", noteId)

	if err := response.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "not found", "message": fmt.Sprintf("Note %v does not exist", noteId)})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": note})

}
