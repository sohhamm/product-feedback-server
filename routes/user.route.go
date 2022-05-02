package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sohhamm/product-feedback-server/controllers"
)

func SetupOAuthRoutesV1(router fiber.Router) {
	router.Get("/auth/:provider", controllers.Login)
	router.Get("/auth/:provider/callback", controllers.GoogleCallback)
	router.Get("/auth/:provider/logout", controllers.GoogleLogout)
}
