package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"ui/handlers"
)

func main() {
	app := fiber.New()
	handlers.InitRoutes(app)

	err := app.Listen(":8082")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
