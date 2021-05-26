package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/woodencoder/startern-go/core/database"
	"github.com/woodencoder/startern-go/core/routes"
)
func main() {
	database.Connect()

	app := fiber.New()

	// Authentication with HttpOnly
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}