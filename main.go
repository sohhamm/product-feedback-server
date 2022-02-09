package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sohhamm/product-feedback-server/database"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the database
	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Product Feedback Server")
	})

	// Health route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Server is running")
	})

	app.Listen(":9000")
}
