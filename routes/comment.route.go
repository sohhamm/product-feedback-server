package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sohhamm/product-feedback-server/controllers"
)

func CommentRoutes(router fiber.Router) {
	comment := router.Group("/comments")

	comment.Get("/", controllers.GetAllFeedbacks)
	comment.Post("/", controllers.CreateFeedback)
	comment.Get("/:id", controllers.GetFeedbackByID)
	comment.Patch("/:id", controllers.UpdateFeedbackByID)
	comment.Delete("/:id", controllers.DeleteFeedbackByID)

}
