package main

import (
	"log"

	"my_hero_server/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	// Register routes
	routes.SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
