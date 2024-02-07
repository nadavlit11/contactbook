package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func InitControllers(port *string) {

	InitServices()

	app := fiber.New(fiber.Config{})

	usersApis := app.Group("/users")
	usersApis.Post("/register/:name", Register)

	contactsApis := app.Group("/contacts")
	contactsApis.Get("/:page", GetContacts)
	contactsApis.Post("/add", CreateContact)
	contactsApis.Get("/search", Search)
	contactsApis.Put("/edit", Edit)
	contactsApis.Delete("/delete/:id", Delete)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}
