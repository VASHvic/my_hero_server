package routes

import (
	"my_hero_server/handlers"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api") // Grouping API routes

	user := api.Group("/users")
	user.Get("/", handlers.GetUsers)
	user.Post("/", handlers.CreateUser)
}
