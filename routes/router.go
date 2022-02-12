package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	v1 := api.Group("/v1")

	FeedbackRoutes(v1)
	CommentRoutes(v1)
	AuthRoutes(v1)
	UserRoutes(v1)

}
