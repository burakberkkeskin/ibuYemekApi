package main

import (
	"ibu-yemek-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.LunchRoute(app)
	routes.HealtchCheckRoute(app)

	app.Listen(":3000")

}
