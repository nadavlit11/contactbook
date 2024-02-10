package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
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

func Login(c *fiber.Ctx) error {
	userId := c.Get("user_id")
	err := usersService.Login(cast.ToInt(userId))

	success := true
	if err != nil {
		success = false
	}

	return c.JSON(fiber.Map{
		"success": success,
	})
}
