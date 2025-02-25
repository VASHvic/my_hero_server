package routes

import (
	"my_hero_server/handlers"

	"github.com/gofiber/fiber/v3"
)

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api")

	user := api.Group("/users")
	user.Get("/", handlers.GetUsers)
	user.Post("/", handlers.CreateUser)
}
