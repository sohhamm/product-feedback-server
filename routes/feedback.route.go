package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sohhamm/product-feedback-server/controllers"
)

func SetupFeedbackRoutesV1(router fiber.Router) {
	feedback := router.Group("/feedbacks")

	feedback.Get("/", controllers.GetAllFeedbacks)
	feedback.Post("/", controllers.CreateFeedback)
	feedback.Get("/:id", controllers.GetFeedbackByID)
	feedback.Patch("/:id", controllers.UpdateFeedbackByID)
	feedback.Delete("/:id", controllers.DeleteFeedbackByID)

}
