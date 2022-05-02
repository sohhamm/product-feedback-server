package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	v1 := api.Group("/v1")

	SetupOAuthRoutesV1(v1)
	SetupFeedbackRoutesV1(v1)
	SetupCommentRoutesV1(v1)
}
