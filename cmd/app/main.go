package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minkgkyaw9899/go-ticket-booking-api/internal/handlers"
	"github.com/minkgkyaw9899/go-ticket-booking-api/internal/repositories"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Go Ticket Booking API",
	})

	server := app.Group("/api")

	// repositories
	eventRepository := repositories.NewEventRepository(nil)

	// handlers
	handlers.NewEventHandler(server.Group("/events"), eventRepository)

	app.Listen(":5000")
}
