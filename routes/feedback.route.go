package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sohhamm/product-feedback-server/controllers"
)

func FeedbackRoutes(router fiber.Router) {
	feedback := router.Group("/feedbacks")

	feedback.Get("/", controllers.GetAllFeedbacks)
	feedback.Post("/", controllers.CreateFeedback)
	feedback.Get("/:id", controllers.GetFeedbackByID)
	feedback.Patch("/:id", controllers.UpdateFeedbackByID)

}
