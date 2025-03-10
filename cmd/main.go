package main

import (


	"github.com/Asker231/Jobber.git/internal/home"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	home.NewHomeHandler(app)

	app.Listen(":3003")
}