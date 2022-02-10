package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sohhamm/product-feedback-server/database"
	"github.com/sohhamm/product-feedback-server/models"
)

func GetAllFeedbacks(c *fiber.Ctx) error {

	db := database.DB.Db
	feedbacks := []models.Feedback{}

	db.Find(&feedbacks)

	return c.Status(fiber.StatusOK).JSON(feedbacks)

}

func CreateFeedback(c *fiber.Ctx) error {
	// db := database.DB

	feedback := new(models.Feedback)

	// var feedback model.Feedback

	if err := c.BodyParser(feedback); err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}

	return c.Status(fiber.StatusOK).JSON(feedback)

}
