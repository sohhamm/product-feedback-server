package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/sohhamm/product-feedback-server/database"
	"github.com/sohhamm/product-feedback-server/models"
	"gorm.io/gorm"
)

func GetAllFeedbacks(c *fiber.Ctx) error {
	feedbacks := []models.Feedback{}

	database.DB.Db.Find(&feedbacks)

	return c.Status(fiber.StatusOK).JSON(feedbacks)

}

func CreateFeedback(c *fiber.Ctx) error {
	feedback := new(models.Feedback)

	if err := c.BodyParser(feedback); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	result := database.DB.Db.Create(&feedback)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error)
	}

	return c.Status(fiber.StatusOK).JSON(feedback)

}

func GetFeedbackByID(c *fiber.Ctx) error {

	id := c.Params("id")

	var feedback models.Feedback

	result := database.DB.Db.First(&feedback, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "record not found"})
	}

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error)
	}

	return c.Status(fiber.StatusOK).JSON(feedback)

}

func UpdateFeedbackByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var feedback models.Feedback

	req := new(models.Feedback)

	result := database.DB.Db.First(&feedback, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "record not found"})
	}

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error)
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if len(req.Title) > 0 {
		feedback.Title = req.Title
	}

	if len(req.Category) > 0 {
		feedback.Category = req.Category
	}

	if len(req.Description) > 0 {
		feedback.Description = req.Description
	}

	if len(req.Status) > 0 {
		feedback.Status = req.Status
	}

	save := database.DB.Db.Save(&feedback)

	if save.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error)
	}

	return c.Status(fiber.StatusOK).JSON(feedback)
}
