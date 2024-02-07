package main

import (
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	name := c.Params("name")
	userId, err := usersService.CreateUser(name)

	success := true
	if err != nil {
		success = false
	}

	return c.JSON(fiber.Map{
		"success": success,
		"userId":  userId,
	})
}
