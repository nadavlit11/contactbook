package main

import (
	"contactbook/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/cast"
)

func GetContacts(c *fiber.Ctx) error {
	userId := cast.ToInt(c.Get("user_id"))
	if userId == 0 {
		return c.JSON(fiber.Map{
			"success": false,
		})
	}

	page, err := c.ParamsInt("page", 1)
	if err != nil {
		log.Error(err)
	}

	contacts, err := contactsService.GetContactsPage(userId, page, pageSize)

	success := true
	if err != nil {
		success = false
	}

	return c.JSON(fiber.Map{
		"success":  success,
		"contacts": contacts,
	})
}

func CreateContact(c *fiber.Ctx) error {
	userId := cast.ToInt(c.Get("user_id"))

	body := c.Body()
	var contact models.Contact
	err := json.Unmarshal(body, &contact)
	if err != nil {
		log.Error("bad input")
		return c.JSON(fiber.Map{
			"success": false,
		})
	}

	err = contactsService.InsertContact(userId, contact)
	success := true
	if err != nil {
		success = false
	}

	return c.JSON(fiber.Map{
		"success": success,
	})
}

func Search(c *fiber.Ctx) error {
	userId := cast.ToInt(c.Get("user_id"))

	body := c.Body()
	var contact models.Contact
	err := json.Unmarshal(body, &contact)
	if err != nil {
		log.Error("bad input")
		return c.JSON(fiber.Map{
			"success": false,
		})
	}

	users, err := contactsService.Search(userId, contact)

	success := true
	if err != nil {
		success = false
	}

	return c.JSON(fiber.Map{
		"success": success,
		"users":   users,
	})
}

func Edit(c *fiber.Ctx) error {
	userId := cast.ToInt(c.Get("user_id"))

	body := c.Body()
	var contact models.Contact
	err := json.Unmarshal(body, &contact)
	if err != nil {
		log.Error("bad input")
		return c.JSON(fiber.Map{
			"success": false,
		})
	}

	err = contactsService.Edit(userId, contact)

	success := true
	if err != nil {
		success = false
	}

	return c.JSON(fiber.Map{
		"success": success,
	})
}

func Delete(c *fiber.Ctx) error {
	userId := cast.ToInt(c.Get("user_id"))

	id, err := c.ParamsInt("id")
	if err != nil {
		log.Error(err)
	}
	err = contactsService.Delete(userId, id)

	success := true
	if err != nil {
		success = false
	}

	return c.JSON(fiber.Map{
		"success": success,
	})
}
