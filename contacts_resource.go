package main

import (
	"contactbook/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func GetContacts(c *fiber.Ctx) error {
	page, err := c.ParamsInt("page", 1)
	if err != nil {
		log.Error(err)
	}
	users, err := contactsService.GetContactsPage(page, pageSize)

	success := true
	if err != nil {
		success = false
	}

	return c.JSON(fiber.Map{
		"success": success,
		"users":   users,
	})
}

func CreateContact(c *fiber.Ctx) error {
	body := c.Body()
	var contact models.Contact
	err := json.Unmarshal(body, &contact)
	if err != nil {
		log.Error("bad input")
		return c.JSON(fiber.Map{
			"success": false,
		})
	}

	err = contactsService.InsertContact(contact)
	success := true
	if err != nil {
		success = false
	}

	return c.JSON(fiber.Map{
		"success": success,
	})
}

func Search(c *fiber.Ctx) error {
	body := c.Body()
	var contact models.Contact
	err := json.Unmarshal(body, &contact)
	if err != nil {
		log.Error("bad input")
		return c.JSON(fiber.Map{
			"success": false,
		})
	}

	users, err := contactsService.Search(contact)

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
	body := c.Body()
	var contact models.Contact
	err := json.Unmarshal(body, &contact)
	if err != nil {
		log.Error("bad input")
		return c.JSON(fiber.Map{
			"success": false,
		})
	}

	err = contactsService.Edit(contact)

	success := true
	if err != nil {
		success = false
	}

	return c.JSON(fiber.Map{
		"success": success,
	})
}

func Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Error(err)
	}
	err = contactsService.Delete(id)

	success := true
	if err != nil {
		success = false
	}

	return c.JSON(fiber.Map{
		"success": success,
	})
}
