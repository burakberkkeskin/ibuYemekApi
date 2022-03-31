package routes

import (
	"ibu-yemek-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func HealtchCheckRoute(app *fiber.App) {

	app.Get("/healthcheck", controllers.Healthcheck)
}
