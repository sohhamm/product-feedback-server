package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sohhamm/product-feedback-server/config"
	"github.com/sohhamm/product-feedback-server/database"
)

func main() {
	// Connect to the database
	database.ConnectDB()

	// Start a new fiber app
	app := fiber.New()

	// Cors middleware
	app.Use(cors.New())

	// Health route
	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.SendString("Server is running")
	})

	//Set up routes

	// Not found route
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "route not found"})
	})

	log.Fatal(app.Listen(config.Config("SERVER_PORT")))
}
