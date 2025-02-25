package routes

import (
	"my_hero_server/handlers"

	"github.com/gofiber/fiber/v3"
)

func SetupQuestRoutes(app *fiber.App) {
	api := app.Group("/api")

	quests := api.Group("/quests")
	quests.Get("/", handlers.GetQuests)
	quests.Get("/random", handlers.GetRandom)
}
