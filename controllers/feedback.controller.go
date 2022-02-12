package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sohhamm/product-feedback-server/database"
	"github.com/sohhamm/product-feedback-server/models"
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
