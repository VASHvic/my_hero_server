package main

import (
	"log"

	"my_hero_server/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	// Register routes
	routes.SetupUserRoutes(app)
	routes.SetupQuestRoutes(app)

	// Start server
	// log.Fatal(app.Listen("127.0.0.1:3000"))
	log.Fatal(app.Listen(":3000"))
}
