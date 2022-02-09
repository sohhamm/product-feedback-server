package router

import "github.com/gofiber/fiber/v2"

func FeedbackRoutes(router fiber.Router) {
	feedback := router.Group("/feedback")

	feedback.Get("/feedbacks")

}
