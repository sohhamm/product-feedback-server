package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/sohhamm/product-feedback-server/config"
	"github.com/sohhamm/product-feedback-server/database"
	"github.com/sohhamm/product-feedback-server/routes"
)

func main() {
	// Connect to the database
	database.ConnectDB()

	goth.UseProviders(google.New(config.Config("GOOGLE_CLIENT_ID"), config.Config("GOOGLE_CLIENT_SECRET"), "http://localhost:9000/api/v1/auth/google/callback"))
	// Start a new fiber app
	app := fiber.New()
	// Cors middleware
	app.Use(cors.New())
	// Health route
	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.SendString("Server is running")
	})

	//Set up routes
	routes.SetupRoutes(app)
	app.Get("/test", func(ctx *fiber.Ctx) error {
		ctx.Format("<p><a href='api/v1/auth/google'>google</a></p>")
		return nil
	})
	// Not found route
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "route not found"})
	})
	log.Fatal(app.Listen(config.Config("SERVER_PORT")))
}
