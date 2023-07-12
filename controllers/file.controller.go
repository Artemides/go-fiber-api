package controllers

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

const DOMAIN = "http://localhost:4000/images/"

func UploadFilesController(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}
	files := form.File["images"]
	filePaths := []string{}
	for _, file := range files {
		filename := file.Filename
		fileExt := filepath.Ext(filename)
		originalName := strings.TrimSuffix(filepath.Base(filename), fileExt)
		now := time.Now()
		fileName := strings.ReplaceAll(strings.ToLower(originalName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
		filePath := DOMAIN + fileName
		err := c.SaveFile(file, fmt.Sprintf("./public/images/%s", fileName))
		if err != nil {
			log.Fatal(err)
		}
		filePaths = append(filePaths, filePath)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"filePaths": filePaths})

}

// func uploadSingleFile(ctx *gin.Context) {
// 	file, header, err := ctx.Request.FormFile("image")
// 	if err != nil {
// 		ctx.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
// 		return
// 	}

// 	fileExt := filepath.Ext(header.Filename)
// 	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
// 	now := time.Now()
// 	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
// 	filePath := "http://localhost:8000/images/single/" + filename

// 	out, err := os.Create("public/single/" + filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer out.Close()
// 	_, err = io.Copy(out, file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{"filepath": filePath})
// }
