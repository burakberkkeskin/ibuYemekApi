package controllers

import (
	"ibu-yemek-api/cache"
	"ibu-yemek-api/responses"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetLunch(c *fiber.Ctx) error {

	log.Println("GetLaunch Request For", c.Params("day"))
	day := c.Params("day")

	if cache.CheckDate(&day) == false {
		response := responses.LunchResponse{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Data:    &fiber.Map{"data": "Wrong Day"},
		}
		return c.Status(http.StatusBadRequest).JSON(&response)
	} else if cache.IsEmpty(c.Params("day")) {
		response := responses.LunchResponse{
			Status:  http.StatusNotFound,
			Message: "Error",
			Data:    &fiber.Map{"data": "Yemek Bulunamadı"},
		}
		return c.Status(http.StatusBadRequest).JSON(&response)
	} else {
		response := responses.LunchResponse{
			Status:  http.StatusOK,
			Message: "Success",
			Data:    cache.GetLunch(&day),
		}
		return c.Status(http.StatusOK).JSON(&response)
	}

}
