package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sohhamm/product-feedback-server/database"
	"github.com/sohhamm/product-feedback-server/model"
)

func GetAllFeedbacks(c *fiber.Ctx) {

	db := database.DB
	var feedbacks []model.Feedback

	db.Find(&feedbacks)

	c.Status(200).JSON(feedbacks)

}
