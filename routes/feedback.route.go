package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sohhamm/product-feedback-server/controllers"
)

func FeedbackRoutes(router fiber.Router) {
	feedback := router.Group("/feedback")

	// Get all feedbacks
	feedback.Get("/", controllers.GetAllFeedbacks)

}
