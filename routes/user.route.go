package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sohhamm/product-feedback-server/controllers"
)

func UserRoutes(router fiber.Router) {
	router.Get("/me", controllers.CreateFeedback)

}

func AuthRoutes(router fiber.Router) {

	router.Post("/register", controllers.CreateFeedback)
	router.Get("/login", controllers.CreateFeedback)

}
