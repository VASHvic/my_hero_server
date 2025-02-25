package handlers

import (
	"fmt"
	"math/rand"
	"my_hero_server/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

var quests = []models.Quest{
	{
		ID:          1,
		Title:       "Trial of Strength",
		Description: "Do a weightlifting session",
		Stat:        "strength",
	},
	{
		ID:          2,
		Title:       "Arcane Studies",
		Description: "Expand your knowledge through reading, or studying",
		Stat:        "intelligence",
	},
	{
		ID:          3,
		Title:       "Diplomatic Mission",
		Description: "Engage in meaningful conversation",
		Stat:        "charisma",
	},
	{
		ID:          4,
		Title:       "Meditation",
		Description: "Reflect on your experiences",
		Stat:        "wisdom",
	},
	{
		ID:          5,
		Title:       "Precision Throw",
		Description: "Play a game of darts or practice throwing at a target",
		Stat:        "dexterity",
	},
	{
		ID:          6,
		Title:       "Endurance Run",
		Description: "Complete a cardio session such as running or cycling",
		Stat:        "constitution",
	},
}

func GetQuests(c fiber.Ctx) error {
	return c.JSON(quests)
}

func GetRandom(c fiber.Ctx) error {
	nStr := c.Query("n", "1")
	n, err := strconv.Atoi(nStr)
	if err != nil {
		fmt.Printf("Error parsing n: %v\n", err)
		n = 1
	}
	fmt.Printf("Requested %d quests\n", n)

	if n > len(quests) {
		fmt.Printf("Adjusting n from %d to %d (max available)\n", n, len(quests))
		n = len(quests)
	}

	// Create a new slice to avoid modifying the original
	shuffledQuests := make([]models.Quest, len(quests))
	copy(shuffledQuests, quests)

	// Shuffle the quests
	rand.Shuffle(len(shuffledQuests), func(i, j int) {
		shuffledQuests[i], shuffledQuests[j] = shuffledQuests[j], shuffledQuests[i]
	})

	return c.JSON(shuffledQuests[:n])
}
