package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func InitControllers(port *string) {

	InitServices()

	app := fiber.New(fiber.Config{})

	// Bind handlers
	app.Get("/contacts/:page", GetContacts)
	app.Post("/contact", CreateContact)
	app.Get("/search", Search)
	app.Put("/edit", Edit)
	app.Delete("/contact/:id", Delete)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}
