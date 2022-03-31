package routes

import (
	"ibu-yemek-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func LunchRoute(app *fiber.App) {

	app.Get("/day/:day", controllers.GetLunch)

}
