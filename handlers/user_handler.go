package handlers

import (
	"github.com/gofiber/fiber/v3"
)

func GetUsers(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "List of users...",
	})
}

func CreateUser(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "User created",
	})
}
